package stringset

// Union returns a new StringSet which is the union of all provided sets.
func Union(sets ...StringSet) StringSet {
	out := make(StringSet)
	for _, set := range sets {
		out = out.union(set)
	}
	return out
}

// Union returns a new StringSet which is the union of all provided sets.
func (ss StringSet) Union(sets ...StringSet) StringSet {
	return Union(append(sets, ss)...)
}

func (ss StringSet) union(set StringSet) StringSet {
	out := ss.Copy()
	for str := range set {
		out.add(str)
	}
	return out
}
