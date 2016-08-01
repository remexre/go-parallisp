package stringset

// Intersection returns a new StringSet which is the intersection of all
// provided sets.
func Intersection(sets ...StringSet) StringSet {
	out := make(StringSet)
	for _, set := range sets {
		out = out.intersection(set)
	}
	return out
}

// Intersection returns a new StringSet which is the intersection of all
// provided sets.
func (ss StringSet) Intersection(sets ...StringSet) StringSet {
	return Intersection(append(sets, ss)...)
}

func (ss StringSet) intersection(set StringSet) StringSet {
	out := New()
	for str := range set {
		if ss.Contains(str) {
			out.add(str)
		}
	}
	return out
}
