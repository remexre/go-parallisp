package exprset

// Intersection returns a new ExprSet which is the intersection of all
// provided sets.
func Intersection(sets ...ExprSet) ExprSet {
	out := make(ExprSet)
	for _, set := range sets {
		out = out.intersection(set)
	}
	return out
}

// Intersection returns a new ExprSet which is the intersection of all
// provided sets.
func (es ExprSet) Intersection(sets ...ExprSet) ExprSet {
	return Intersection(append(sets, es)...)
}

func (es ExprSet) intersection(set ExprSet) ExprSet {
	out := New()
	for str := range set {
		if es.Contains(str) {
			out.add(str)
		}
	}
	return out
}
