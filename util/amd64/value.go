package amd64

// Outputtable is an interface for "things" which can be outputted in GNU or
// Intel syntax.
type Outputtable interface {
	// GnuString converts the function or value to a GNU-syntax string.
	GnuString() string
	// IntelString converts the function or value to an Intel-syntax string.
	IntelString() string
}

// Value is an interface for assembly values.
type Value interface {
	Outputtable
	Amd64Value() // Tag function
}
