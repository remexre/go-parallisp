package types

// A Slicer is a type that can be "sliced."
type Slicer interface {
	Get(Integer) (Expr, error)
	Slice(from, to Integer) (Expr, error)
}
