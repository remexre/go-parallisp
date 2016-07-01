package main

import (
	"fmt"

	"remexre.xyz/parallisp/parser"
)

const src = `

(defun f [x]
	(+ (^ x 2) x 1))
(f 12)

`

func main() {
	exprs, err := parser.ParseAll(src)
	if err != nil {
		panic(err)
	}

	for _, expr := range exprs {
		fmt.Println(expr.Expr())
	}
}
