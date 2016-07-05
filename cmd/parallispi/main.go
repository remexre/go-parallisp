package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"remexre.xyz/parallisp/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "FILE")
		return
	}

	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	exprs, err := parser.Parse(string(src))
	if err != nil {
		panic(err)
	}

	b, err := xml.MarshalIndent(exprs, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
