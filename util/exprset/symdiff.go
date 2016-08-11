package exprset

// SymDifference returns a new ExprSet which is the symmetric difference of
// all provided sets.
func SymDifference(sets ...ExprSet) ExprSet {
	out := make(ExprSet)
	for _, set := range sets {
		out = out.symDifference(set)
	}
	return out
}

// SymDifference returns a new ExprSet which is the symmetric difference of
// all provided sets.
func (es ExprSet) SymDifference(sets ...ExprSet) ExprSet {
	return SymDifference(append(sets, es)...)
}

func (es ExprSet) symDifference(set ExprSet) ExprSet {
	out := New()
	for str := range set {
		if !es.Contains(str) {
			out.add(str)
		}
	}
	for str := range es {
		if !set.Contains(str) {
			out.add(str)
		}
	}
	return out
}
