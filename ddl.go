package ddl

func If[T any](condition bool, thenValue, elseValue T) T {
	if condition {
		return thenValue
	}

	return elseValue
}
