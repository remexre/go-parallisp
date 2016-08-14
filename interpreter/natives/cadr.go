package natives

import "github.com/remexre/go-parallisp/types"

// Car returns the car of a cons.
func Car(cons types.Cons) types.Expr {
	return cons.Car()
}

// Cdr returns the cdr of a cons.
func Cdr(cons types.Cons) types.Expr {
	return cons.Cdr()
}
