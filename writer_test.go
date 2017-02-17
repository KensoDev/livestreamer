package livestreamer

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestWriter(t *testing.T) { TestingT(t) }

type WriterSuite struct{}

var _ = Suite(&WriterSuite{})

func (s *WriterSuite) TestGenerateFileName(c *C) {
	timeString := GenerateFileName("02-01-2006")
	c.Assert(timeString, Equals, "16-02-2017")
}

func (s *WriterSuite) TestFileWrite(c *C) {
	filename := "fixtures/some-file-name"
	_ = os.Remove(filename)

	stream := &LiveStream{
		Title:       "Some Title",
		Description: "Some Description",
	}

	writer := NewWriter(filename)
	writer.WriteFile(stream)

	file, err := os.Open(filename)
	c.Assert(err, IsNil)
	fileInfo, err := file.Stat()
	c.Assert(err, IsNil)
	c.Assert(fileInfo.Name(), Equals, "some-file-name")
}
