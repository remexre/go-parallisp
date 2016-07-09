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

// SetPrintFilter sets the regexp used to determine whether a message should be
// logged or not.
func SetPrintFilter(regex string) error {
	re, err := regexp.Compile(regex)
	if err != nil {
		return err
	}
	printRegexp.Store(re)
	return nil
}
