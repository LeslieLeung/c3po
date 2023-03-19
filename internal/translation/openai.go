package translation

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type OpenAI struct {
	client *openai.Client
	once   sync.Once
}

func (o *OpenAI) getClient() *openai.Client {
	o.once.Do(func() {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			homeDir, _ := os.UserHomeDir()
			if _, err := os.Stat(filepath.Join(homeDir, ".c3pocfg")); err == nil {
				// read from file
				apiKeyBytes, _ := os.ReadFile(filepath.Join(homeDir, ".c3pocfg"))
				apiKey = string(apiKeyBytes)
			}
		}
		o.client = openai.NewClient(apiKey)
	})
	return o.client
}

func (o *OpenAI) createChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := o.getClient().CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo0301,
		Messages: messages,
	})
	if err != nil {
		return "", err
	}
	logrus.Debugf("Token Usage [Prompt: %d, Completion: %d, Total: %d]",
		resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	return resp.Choices[0].Message.Content, nil
}

func (o *OpenAI) CreateTranslation(ctx context.Context, text string, lang string) (string, error) {
	return o.createChatCompletion(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("Translate a i18n locale string to %s. Consider the context of the value to make better translation. Print only the translated string.", lang),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		},
	})
}

func (o *OpenAI) BatchCreateTranslation(ctx context.Context, text string, lang string) (string, error) {
	return o.createChatCompletion(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("Translate a i18n locale string to following languages %s. Consider the context of the value to make better translation. Each line starts with the language name(eg. \"it: ciao, mondo\"), devide each translation with line break(LF), do not add extra format.", lang),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		},
	})
}
