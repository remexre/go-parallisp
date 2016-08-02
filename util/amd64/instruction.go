package amd64

import "strings"

// Instruction is a structure for assembly instructions.
type Instruction struct {
	Opcode Opcode
	Params []Value
}

// NewInstruction creates a new instruction.
func NewInstruction(opcode Opcode, params ...Value) *Instruction {
	return &Instruction{opcode, params}
}

// GnuString converts the function or value to a GNU-syntax string.
func (ins *Instruction) GnuString() string {
	var paramStr string
	if len(ins.Params) > 0 {
		params := make([]string, len(ins.Params))
		l := len(ins.Params) - 1
		for i, param := range ins.Params {
			params[l-i] = param.GnuString()
		}
		paramStr = " " + strings.Join(params, ", ")
	}
	return ins.Opcode.String() + paramStr
}

// IntelString converts the function or value to an Intel-syntax string.
func (ins *Instruction) IntelString() string {
	var paramStr string
	if len(ins.Params) > 0 {
		params := make([]string, len(ins.Params))
		for i, param := range ins.Params {
			params[i] = param.IntelString()
		}
		paramStr = " " + strings.Join(params, ", ")
	}
	return ins.Opcode.String() + paramStr
}
