package livestreamer

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestReader(t *testing.T) { TestingT(t) }

type ReaderSuite struct{}

var _ = Suite(&ReaderSuite{})

func (s *ReaderSuite) TestYamlRead(c *C) {
	reader, err := NewReader("fixtures/test.yml")

	c.Assert(err, IsNil)
	c.Assert(len(reader.FileContent), Not(Equals), 0)
}

func (s *ReaderSuite) TestParsing(c *C) {
	reader, err := NewReader("fixtures/test.yml")
	c.Assert(err, IsNil)

	stream, err := reader.Parse()
	c.Assert(err, IsNil)
	c.Assert(stream.Title, Equals, "This is some title")
	c.Assert(stream.Description, Equals, "This is some description")
	c.Assert(len(stream.Tags), Equals, 3)
	c.Assert(stream.Tags[0], Equals, "A")
}
