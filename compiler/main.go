package compiler

import (
	"fmt"
	"strings"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/util/amd64"
)

// Compile compiles a module.
func Compile(module *ast.Module) (string, error) {
	ch := make(chan amd64.Outputtable, 10)
	errCh := make(chan error)
	go compile(ch, errCh, module)

	out := []amd64.Outputtable{amd64.Comment(module.Name)}
	var mainInstructions []amd64.Instruction
	for o := range ch {
		if ins, ok := o.(*amd64.Instruction); ok {
			mainInstructions = append(mainInstructions, *ins)
		} else if fn, ok := o.(*amd64.Function); ok {
			out = append(out, fn)
		} else {
			panic("unknown return from compileNode")
		}
	}
	out = append(out, &amd64.Function{
		Name: "parallisp$main",
		Contents: append(
			mainInstructions,
			*amd64.NewInstruction(amd64.RET),
		),
	})
	for err := range errCh {
		if err != nil {
			return "", err
		}
	}

	outStrs := make([]string, len(out))
	for i, o := range out {
		// TODO Configure Gnu vs Intel
		outStrs[i] = o.IntelString()
	}
	return strings.Join(outStrs, "\n\n"), nil
}

func compile(ch chan<- amd64.Outputtable, errCh chan<- error, module *ast.Module) {
	defer close(ch)
	defer close(errCh)

	for _, importNode := range module.Imports {
		errCh <- fmt.Errorf("TODO Import %v", importNode)
		return
	}
	for _, node := range module.Body {
		if err := compileNode(ch, node); err != nil {
			errCh <- err
			return
		}
	}
}

func compileNode(ch chan<- amd64.Outputtable, node ast.Node) error {
	switch n := node.(type) {
	case ast.Integer:
		ch <- amd64.NewInstruction(
			amd64.MOV,
			amd64.Rax,
			amd64.Immediate(n),
		)
	default:
		return fmt.Errorf("compiler.Compile: cannot compile %s: unknown type", node)
	}
	return nil
}
