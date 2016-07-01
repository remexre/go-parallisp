package number

import "remexre.xyz/go-parcom"

// Parse parses number literals recognized by parallisp.
func Parse(in string) (string, interface{}, bool) {
	return parcom.Map(parcom.Chain(
		sign,
		parcom.Alt(
			decimalNumber,
		),
	), applySign)(in)
}
