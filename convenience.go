package golib

// Adapts a Variable Argument list of primitives to an array of primitives.
// Can also be used to create an empty array of primitives.
func VA(primitives ...Primitive) []Primitive {
	return primitives
}

// Adapts a Variable Argument list of primitive rows to an array of primitive rows.
// Can also be used to create an empty array of primtive rows.
func VAA(rows ...[]Primitive) [][]Primitive {
	return rows
}
