package cmd

import (
	"context"
	"fmt"
	"github.com/leslieleung/c3po/internal/file"
	"github.com/leslieleung/c3po/internal/translation"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var translateFileCmd = &cobra.Command{
	Use:    "translateFile",
	PreRun: toggleDebug,
	Short:  "Translate i18n file",
	Run: func(cmd *cobra.Command, args []string) {
		translateFile()
	},
}

var (
	inFileType  string
	outFileType string
	fileName    string
)

func translateFile() {
	openai := translation.OpenAI{}

	inFileType = filepath.Ext(fileName)[1:]

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}()
	inHandler := file.GetHandler(inFileType)
	keys, texts, headers, err := inHandler.Read(f)
	if err != nil {
		panic(err)
	}
	// remove key and the first language
	languageHeaders := headers[2:]
	languages := strings.Join(languageHeaders, ",")
	out := make([][]string, len(texts))

	ctx := context.Background()

	for i, s := range texts {
		idx := i
		oriStr := s
		ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		trans, err := openai.BatchCreateTranslation(ctxTimeout, oriStr, languages)
		if err != nil {
			fmt.Println(err.Error())
		}
		translations := parseTranslation(trans)
		out[idx] = append([]string{keys[idx], texts[idx]}, translations...)
		logrus.Debugf("out[%v]: %v", idx, out[idx])
	}
	logrus.Debugf("out: %v", out)

	outFile, _ := os.OpenFile(parseOutputFileName(fileName), os.O_RDWR|os.O_CREATE, 0755)
	outHandler := file.GetHandler(outFileType)
	err = outHandler.Write(outFile, headers, out)
	if err != nil {
		panic(err)
	}
	outFile.Close()
}

func parseTranslation(trans string) []string {
	lines := strings.Split(trans, "\n")
	logrus.Debugf("lines: %v", lines)
	translations := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < 4 {
			logrus.Errorf("OpenAI returned a malformed translation, please try again: %v", line)
			panic("OpenAI returned a malformed translation, please try again")
		}
		translations[i] = line[4:]
	}
	return translations
}

func parseOutputFileName(filename string) string {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(ext)] + "_translation" + ext
}

func init() {
	rootCmd.AddCommand(translateFileCmd)

	translateFileCmd.Flags().StringVarP(&fileName, "fileName", "f", "", "file name")
	translateFileCmd.Flags().StringVarP(&outFileType, "outFileType", "o", "csv", "output file type")
}
