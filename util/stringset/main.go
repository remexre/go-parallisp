package stringset

import (
	"strconv"
	"strings"
)

// StringSet is a type for a set of strings.
type StringSet map[string]struct{}

// New returns a new StringSet with the given elements.
func New(strs ...string) StringSet {
	if len(strs) == 0 {
		return make(StringSet)
	}
	return make(StringSet, len(strs)).Add(strs...)
}

// Copy returns a copy of a StringSet.
func (ss StringSet) Copy() StringSet {
	out := make(StringSet)
	for str := range ss {
		out.add(str)
	}
	return out
}

// Contains returns whether StringSet contains the given string.
func (ss StringSet) Contains(str string) bool {
	_, ok := ss[str]
	return ok
}

// Equals returns whether a set is equal to another.
func (ss StringSet) Equals(other StringSet) bool {
	if ss.Length() != other.Length() {
		return false
	}
	return ss.Difference(other).Length() == 0
}

// Length returns the cardinality of the set.
func (ss StringSet) Length() int {
	return len(ss)
}

// String returns the string representation of the set.
func (ss StringSet) String() string {
	strs := make([]string, 0, len(ss))
	for str := range ss {
		strs = append(strs, strconv.Quote(str))
	}

	return "{" + strings.Join(strs, " ") + "}"
}

// ToSlice returns the elements of the set as a slice.
func (ss StringSet) ToSlice() []string {
	var out []string
	for str := range ss {
		out = append(out, str)
	}
	return out
}
