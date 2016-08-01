package stringset

// SymDifference returns a new StringSet which is the symmetric difference of
// all provided sets.
func SymDifference(sets ...StringSet) StringSet {
	out := make(StringSet)
	for _, set := range sets {
		out = out.symDifference(set)
	}
	return out
}

// SymDifference returns a new StringSet which is the symmetric difference of
// all provided sets.
func (ss StringSet) SymDifference(sets ...StringSet) StringSet {
	return SymDifference(append(sets, ss)...)
}

func (ss StringSet) symDifference(set StringSet) StringSet {
	out := New()
	for str := range set {
		if !ss.Contains(str) {
			out.add(str)
		}
	}
	for str := range ss {
		if !set.Contains(str) {
			out.add(str)
		}
	}
	return out
}
