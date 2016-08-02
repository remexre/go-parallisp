package amd64

import "fmt"

// Immediate is a type for immediate values.
type Immediate uint64

// Amd64Value is a tag function for values.
func (Immediate) Amd64Value() {}

// GnuString converts the function or value to a GNU-syntax string.
func (r Immediate) GnuString() string {
	return "$" + fmt.Sprint(r)
}

// IntelString converts the function or value to an Intel-syntax string.
func (r Immediate) IntelString() string {
	return fmt.Sprint(r)
}
