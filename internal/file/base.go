package file

import "os"

type Handler interface {
	Read(file *os.File) ([]string, []string, []string, error)
	Write(file *os.File, headers []string, out [][]string) error
}

func GetHandler(fileType string) Handler {
	switch fileType {
	case "csv":
		return NewCSVHandler()
	default:
		panic(fileType + "not supported")
	}
}
