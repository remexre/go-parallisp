package exprset

// Union returns a new ExprSet which is the union of all provided sets.
func Union(sets ...ExprSet) ExprSet {
	out := make(ExprSet)
	for _, set := range sets {
		out = out.union(set)
	}
	return out
}

// Union returns a new ExprSet which is the union of all provided sets.
func (es ExprSet) Union(sets ...ExprSet) ExprSet {
	return Union(append(sets, es)...)
}

func (es ExprSet) union(set ExprSet) ExprSet {
	out := es.Copy()
	for str := range set {
		out.add(str)
	}
	return out
}
