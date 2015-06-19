package route

import (
	"bytes"
	"regexp"
)

type goRouteBuilder struct {
	buffer bytes.Buffer
}

var start = "^"
var end = "/?$"

// NewBuilder creates a new goRouteBuilder
func NewBuilder() Builder {
	return &goRouteBuilder{buffer: bytes.Buffer{}}
}

func (b *goRouteBuilder) Append(s string) Builder {
	b.buffer.WriteString(s)
	return b
}

func (b *goRouteBuilder) AppendPart(s string) Builder {
	b.Append("/")
	b.Append(s)
	return b
}

func (b *goRouteBuilder) Fork() Builder {
	nRouteBuilder := &goRouteBuilder{buffer: bytes.Buffer{}}
	nRouteBuilder.Append(b.buffer.String())
	return nRouteBuilder
}

func (b *goRouteBuilder) MustCompile() *regexp.Regexp {
	return regexp.MustCompile(start + b.buffer.String() + end)
}

func (b *goRouteBuilder) Compile() (*regexp.Regexp, error) {
	return regexp.Compile(start + b.buffer.String() + end)
}
