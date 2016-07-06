package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"remexre.xyz/go-parallisp/interpreter"
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

	out, err := interpreter.Interpret(string(src))
	if err != nil {
		panic(err)
	} else if out != nil {
		fmt.Println(out)
	}
}
