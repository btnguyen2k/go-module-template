package mymodule

// sample code

// Clone return a copy of an input slice.
func Clone[E any](input []E) []E {
	output := make([]E, len(input))
	copy(output, input)
	return output
}
