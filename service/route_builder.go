package service

import (
	"bytes"
	"regexp"
)

// RouteBuilder is a Regexp builder for servable URL routes.
type RouteBuilder interface {
	Append(string) RouteBuilder
	AppendPart(string) RouteBuilder
	Fork() RouteBuilder
	MustCompile() *regexp.Regexp
	Compile() (*regexp.Regexp, error)
}

type goRouteBuilder struct {
	buffer bytes.Buffer
}

var start = "^"
var end = "/?$"

func (b *goRouteBuilder) Append(s string) RouteBuilder {
	b.buffer.WriteString(s)
	return b
}

func (b *goRouteBuilder) AppendPart(s string) RouteBuilder {
	b.Append("/")
	b.Append(s)
	return b
}

func (b *goRouteBuilder) Fork() RouteBuilder {
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
