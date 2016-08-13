package exprset

// Intersection returns a new ExprSet which is the intersection of all
// provided sets.
func Intersection(sets ...ExprSet) ExprSet {
	var out ExprSet
	for _, set := range sets {
		out = out.intersection(set)
	}
	if len(sets) == 1 {
		return sets[0]
	}
	return sets[0].intersection(Intersection(sets[1:]...))
}

// Intersection returns a new ExprSet which is the intersection of all
// provided sets.
func (es ExprSet) Intersection(sets ...ExprSet) ExprSet {
	return Intersection(append(sets, es)...)
}

func (es ExprSet) intersection(set ExprSet) ExprSet {
	var out ExprSet
	for _, expr := range set {
		if es.Contains(expr) {
			out = append(out, expr)
		}
	}
	return out
}
