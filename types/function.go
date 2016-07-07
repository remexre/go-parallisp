package types

// Function is a callable interface.
type Function interface {
	Expr

	Call(...Expr) (Expr, error)
	MinArgN() int
	MaxArgN() int
}
