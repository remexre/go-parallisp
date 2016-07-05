package types

// Env represents an execution environment.
type Env interface {
	Derive(map[Symbol]Expr) Env

	Get(Symbol) (Expr, bool)
	Def(Symbol, Expr) error
	List(recursive bool) []Symbol
	Set(Symbol, Expr) error
}
