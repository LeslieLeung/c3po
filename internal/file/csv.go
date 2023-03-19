package file

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"os"
)

type CSVHandler struct{}

func NewCSVHandler() Handler {
	return &CSVHandler{}
}

func (h *CSVHandler) Read(file *os.File) ([]string, []string, []string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, err
	}
	headers := records[0]
	keys := make([]string, 0, len(records)-1)
	texts := make([]string, 0, len(records)-1)
	for i, record := range records {
		if i == 0 {
			continue
		}
		keys = append(keys, record[0])
		texts = append(texts, record[1])
	}
	logrus.Debugf("translating total number of keys: %d", len(keys))
	return keys, texts, headers, nil
}

func (h *CSVHandler) Write(file *os.File, headers []string, out [][]string) error {
	writer := csv.NewWriter(file)
	err := writer.Write(headers)
	if err != nil {
		return err
	}
	err = writer.WriteAll(out)
	if err != nil {
		return err
	}
	return nil
}
