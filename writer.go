package livestreamer

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

func GenerateFileName(format string) string {
	now := time.Now()
	return now.Format(format)
}

func GenerateRandomString() string {
	return "random"
}

type Writer struct {
	FileName string
}

func NewWriter(filename string) *Writer {
	return &Writer{
		FileName: filename,
	}
}

func (w *Writer) WriteFile(stream *LiveStream) error {
	content, err := yaml.Marshal(stream)
	if err != nil {
		return err
	}
	writeErr := ioutil.WriteFile(w.FileName, content, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
