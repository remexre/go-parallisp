package stringset

// Add adds a string or strings to a StringSet, then returns the StringSet.
func (ss StringSet) Add(strs ...string) StringSet {
	for _, str := range strs {
		ss.add(str)
	}
	return ss
}

func (ss StringSet) add(str string) {
	ss[str] = struct{}{}
}
