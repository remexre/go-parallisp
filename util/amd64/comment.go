package amd64

import "fmt"

// Comment is a type for comments.
type Comment string

// GnuString converts the function or value to a GNU-syntax string.
func (r Comment) GnuString() string {
	return "// " + fmt.Sprint(r)
}

// IntelString converts the function or value to an Intel-syntax string.
func (r Comment) IntelString() string {
	return "; " + fmt.Sprint(r)
}
