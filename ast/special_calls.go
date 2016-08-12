package ast

import "remexre.xyz/go-parallisp/types"

// SpecialCalls is a map of all forms that create special AST nodes (i.e. not
// function calls).
var SpecialCalls map[types.Symbol]func([]types.Expr) (Node, error)

func init() {
	SpecialCalls = map[types.Symbol]func([]types.Expr) (Node, error){
		"defun":    NewDefun,
		"defmacro": NewDefmacro,
		"import":   NewImport,
		"lambda":   NewLambda,
		"let":      NewLet,
		"let*":     NewSequentialLet,
		"progn":    NewProgn,
		"quote":    NewQuote,
	}
}
