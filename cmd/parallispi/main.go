package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/remexre/go-parallisp/debug"
	"github.com/remexre/go-parallisp/interpreter"
)

func main() {
	flag.Var(debug.Debugger, "debug", "The regexp for debug logs")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: parallispi FLAGS FILE")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		return
	}

	src, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	out, _, err := interpreter.Interpret(string(src), flag.Arg(0))
	if err != nil {
		panic(err)
	} else if out != nil {
		fmt.Println("returned", out)
	}
}
