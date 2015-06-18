package model

import "time"

// BlogPost is the interface for a generic blog post
type BlogPost interface {
	Date() time.Time
	Uuid() string
}
