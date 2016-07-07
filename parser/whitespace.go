package parser

import "remexre.xyz/go-parcom"

var whitespace = parcom.AnyOf(" \n\t")
var optionalWS = parcom.Opt(whitespace, "")
