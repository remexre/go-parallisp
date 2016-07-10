package types

// Env represents an execution environment.
type Env interface {
	Derive(Env) Env

	All(recursive bool) map[Symbol]Expr
	Get(Symbol) (Expr, bool)
	Def(Symbol, Expr) error
	List(recursive bool) []Symbol
	Set(Symbol, Expr) error
}
