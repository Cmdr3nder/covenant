package route

import "regexp"

// Builder is a Regexp builder for servable URL routes.
type Builder interface {
	Append(string) Builder
	AppendPart(string) Builder
	Fork() Builder
	MustCompile() *regexp.Regexp
	Compile() (*regexp.Regexp, error)
}
