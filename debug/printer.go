package debug

import (
	"fmt"
	"regexp"
	"sync/atomic"
)

var printRegexp atomic.Value

func print(ch <-chan D) {
	for d := range ch {
		if printRegexp.Load().(*regexp.Regexp).MatchString(d.Namespace) {
			fmt.Printf("%s: %s\n", d.Namespace, d.Data)
		}
	}
}

// Debugger is a flag.Value for the print filter.
var Debugger = debugger{}

type debugger struct{}

func (debugger) Set(regex string) error {
	re, err := regexp.Compile(regex)
	if err != nil {
		return err
	}
	printRegexp.Store(re)
	return nil
}

func (debugger) String() string {
	return printRegexp.Load().(*regexp.Regexp).String()
}
