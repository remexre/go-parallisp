package amd64

import "fmt"

// Register is a type for registers.
type Register byte

// Registers
const (
	Rax Register = iota
	Rbx
	Rcx
	Rdx
	Rbp
	Rsp
	Rsi
	Rdi
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

// Amd64Value is a tag function for values.
func (Register) Amd64Value() {}

// GnuString converts the function or value to a GNU-syntax string.
func (r Register) GnuString() string {
	return "%" + r.String()
}

// IntelString converts the function or value to an Intel-syntax string.
func (r Register) IntelString() string {
	return r.String()
}

// String returns the name of the register.
func (r Register) String() string {
	name := fmt.Sprint(byte(r))
	switch r {
	case Rax:
		name = "ax"
	case Rbx:
		name = "bx"
	case Rcx:
		name = "cx"
	case Rdx:
		name = "dx"
	case Rbp:
		name = "bp"
	case Rsp:
		name = "sp"
	case Rsi:
		name = "si"
	case Rdi:
		name = "di"
	}
	return "r" + name
}
