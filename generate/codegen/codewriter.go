package codegen

import (
	"fmt"
	"io"
)

// CodeWriter write code
type CodeWriter struct {
	Writer    io.Writer
	indent    int
	IndentStr string
}

// Indent add more indent from now
func (c *CodeWriter) Indent() {
	c.indent++
}

// UnIndent reduce indent from now
func (c *CodeWriter) UnIndent() {
	c.indent--
	if c.indent < 0 {
		panic("indent less than zero")
	}
}

// WriteLine write one line code
func (c *CodeWriter) WriteLineF(format string, values ...any) {
	c.WriteLine(fmt.Sprintf(format, values...))
}

// WriteLine write one line code
func (c *CodeWriter) WriteLine(line string) {
	for i := 0; i < c.indent; i++ {
		c.Writer.Write([]byte(c.IndentStr))
	}
	c.Writer.Write([]byte(line))
	c.Writer.Write([]byte("\n"))
}

// WriteLines write multi lines code
func (c *CodeWriter) WriteLines(lines []string) {
	for _, line := range lines {
		c.WriteLine(line)
	}
}
