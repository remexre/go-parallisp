package types

// A Sequence is a type that can represents a finite sequence.
type Sequence interface {
	Get(Integer) (Expr, error)
	Len() (Integer, error)
	Slice(from, to Integer) (Expr, error)
}
