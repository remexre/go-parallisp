package stringset

// Difference returns a new StringSet which is the difference of all provided
// sets with the original.
func (ss StringSet) Difference(sets ...StringSet) StringSet {
	for _, set := range sets {
		ss = ss.difference(set)
	}
	return ss
}

func (ss StringSet) difference(set StringSet) StringSet {
	out := New()
	for str := range ss {
		if !set.Contains(str) {
			out.add(str)
		}
	}
	return out
}
