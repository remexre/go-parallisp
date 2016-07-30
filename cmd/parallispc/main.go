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

	nodes := make([]ast.Node, len(exprs))
	for i, expr := range exprs {
		nodes[i], err = ast.Convert(expr)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("ast = ")
	pp.Println(nodes)
}
