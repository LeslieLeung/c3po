package cmd

import (
	"context"
	"fmt"
	"github.com/leslieleung/c3po/internal/translation"
	"github.com/spf13/cobra"
	"strings"
)

var translateCmd = &cobra.Command{
	Use:    "translate",
	PreRun: toggleDebug,
	Short:  "Translate text to a language",
	Run: func(cmd *cobra.Command, args []string) {
		translate()
	},
}

var (
	text string
	lang string
)

func translate() {
	openai := translation.OpenAI{}

	var (
		trans string
		err   error
	)
	if strings.Contains(lang, ",") {
		trans, err = openai.BatchCreateTranslation(context.Background(), text, lang)
	} else {
		trans, err = openai.CreateTranslation(context.Background(), text, lang)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(trans)
}

func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringVarP(&text, "text", "t", "", "text to translate")
	translateCmd.Flags().StringVarP(&lang, "lang", "l", "", "language to translate to")
}
