package debug

import "regexp"

func init() {
	ch := make(chan D, 256)
	Chan = ch
	printRegexp.Store(regexp.MustCompile("^$"))
	go print(ch)
}
