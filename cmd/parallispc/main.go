package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"

	"remexre.xyz/go-parallisp/ast"
	"remexre.xyz/go-parallisp/parser"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: parallispc FLAGS FILE")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		return
	}

	src, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	exprs, err := parser.Parse(string(src))
	if err != nil {
		panic(err)
	}

	module, err := ast.ConvertModule(exprs)
	if err != nil {
		panic(err)
	}
	fmt.Print("module = ")
	pp.Println(module)

	fmt.Print("\n\n\n")
	fmt.Printf("defines = %s\n", module.Defines())
	fmt.Printf("freeVars = %s\n", module.FreeVars())

	for _, node := range module.Imports {
		fmt.Println(node.ToExpr())
	}
	for _, node := range module.Body {
		fmt.Println(node.ToExpr())
	}
}
