package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/compiler"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "Usage: parallispc FLAGS IN OUT")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		os.Exit(-1)
	}
	inFile := flag.Arg(0)
	outFile := flag.Arg(1)

	module, err := ast.LoadModule(inFile)
	if err != nil {
		panic(err)
	}

	asm, err := compiler.Compile(module)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outFile, []byte(asm), 0644)
	if err != nil {
		panic(err)
	}
}
