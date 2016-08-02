package amd64

// Opcode is a type for opcodes.
type Opcode string

// Opcodes
const (
	MOV Opcode = "mov"
	RET Opcode = "ret"
)

func (opcode Opcode) String() string {
	return string(opcode)
}
