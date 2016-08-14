package natives

import "github.com/remexre/go-parallisp/types"

// StringToVector converts a string into a vector of Unicode code points.
func StringToVector(str types.String) types.Vector {
	bytes := []byte(str)
	vec := make(types.Vector, len(bytes))
	for i, b := range bytes {
		vec[i] = types.Integer(b)
	}
	return vec
}
