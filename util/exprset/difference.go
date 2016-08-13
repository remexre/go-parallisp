package exprset

// Difference returns a new ExprSet which is the difference of all provided
// sets with the original.
func (es ExprSet) Difference(sets ...ExprSet) ExprSet {
	for _, set := range sets {
		es = es.difference(set)
	}
	return es
}

func (es ExprSet) difference(set ExprSet) ExprSet {
	out := New()
	for _, expr := range es {
		if !set.Contains(expr) {
			out = append(out, expr)
		}
	}
	return out
}
