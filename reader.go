package livestreamer

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Reader struct {
	FileName    string
	FileContent []byte
}

func NewReader(filename string) (*Reader, error) {
	reader := &Reader{
		FileName: filename,
	}

	err := reader.readFileContent()

	if err != nil {
		return reader, err
	}

	return reader, nil
}

func (r *Reader) readFileContent() error {
	bytes, err := ioutil.ReadFile(r.FileName)

	if err != nil {
		return err
	}

	r.FileContent = bytes
	return nil
}

func (r *Reader) Parse() (LiveStream, error) {
	var stream LiveStream

	err := yaml.Unmarshal(r.FileContent, &stream)
	if err != nil {
		return stream, err
	}

	return stream, nil
}
