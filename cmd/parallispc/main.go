package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/compiler"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: parallispc FLAGS FILE")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		return
	}
	inFile := flag.Arg(0)

	module, err := ast.LoadModule(inFile)
	if err != nil {
		panic(err)
	}

	asm, err := compiler.Compile(module)
	if err != nil {
		panic(err)
	}

	outFile := strings.TrimSuffix(inFile, filepath.Ext(filepath.Base(inFile))) + ".asm"
	err = ioutil.WriteFile(outFile, []byte(asm), 0644)
	if err != nil {
		panic(err)
	}
}
