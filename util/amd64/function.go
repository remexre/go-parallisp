package amd64

import (
	"bytes"
	"strings"
)

// Function is an interface for assembly procedures.
type Function struct {
	Name     string
	Contents []Instruction
	Global   bool
}

// GnuString converts the function to a GNU-syntax string.
func (f *Function) GnuString() string {
	buf := new(bytes.Buffer)
	buf.WriteString(".global ")
	buf.WriteString(f.Name)
	buf.WriteByte('\n')
	buf.WriteString(f.Name)
	buf.WriteByte(':')
	for _, ins := range f.Contents {
		for _, line := range strings.Split(ins.GnuString(), "\n") {
			buf.WriteString("\n\t")
			buf.WriteString(line)
		}
	}
	return buf.String()
}

// IntelString converts the function to an Intel-syntax string.
func (f *Function) IntelString() string {
	buf := new(bytes.Buffer)
	buf.WriteString("GLOBAL ")
	buf.WriteString(f.Name)
	buf.WriteByte('\n')
	buf.WriteString(f.Name)
	buf.WriteByte(':')
	for _, ins := range f.Contents {
		for _, line := range strings.Split(ins.IntelString(), "\n") {
			buf.WriteString("\n\t")
			buf.WriteString(line)
		}
	}
	return buf.String()
}
