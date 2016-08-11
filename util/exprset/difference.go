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
	for str := range es {
		if !set.Contains(str) {
			out.add(str)
		}
	}
	return out
}
