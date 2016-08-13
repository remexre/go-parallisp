package exprset

// Union returns a new ExprSet which is the union of all provided sets.
func Union(sets ...ExprSet) ExprSet {
	if len(sets) == 0 {
		return nil
	}
	return sets[0].Union(sets[1:]...)
}

// Union returns a new ExprSet which is the union of all provided sets.
func (es ExprSet) Union(sets ...ExprSet) ExprSet {
	for _, set := range sets {
		es = es.union(set)
	}
	return es
}

func (es ExprSet) union(other ExprSet) ExprSet {
	for _, expr := range other {
		es = append(es, expr)
	}
	return es
}
