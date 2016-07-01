package parser

import "remexre.xyz/go-parcom"

var whitespace = parcom.Opt(parcom.AnyOf(" \n\t"), "")
