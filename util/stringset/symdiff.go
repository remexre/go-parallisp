package stringset

// Difference returns a new StringSet which is the symmetric difference of all
// provided sets.
func Difference(sets ...StringSet) StringSet {
	out := make(StringSet)
	for _, set := range sets {
		out = out.difference(set)
	}
	return out
}

// Difference returns a new StringSet which is the symmetric difference of all
// provided sets.
func (ss StringSet) Difference(sets ...StringSet) StringSet {
	return Difference(append(sets, ss)...)
}

func (ss StringSet) difference(set StringSet) StringSet {
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
