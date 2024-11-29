package assert

func ContainsIgnoringFields[T any](actual []T, expected T, ignoredFields ...string) bool {
	for _, value := range actual {
		if EqualsIgnoringFields(value, expected, ignoredFields...) {
			return true
		}
	}

	return false
}
