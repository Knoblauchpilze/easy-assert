package assert

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func EqualsIgnoringFields[T any](actual T, expected T, ignoredFields ...string) bool {
	var dummy T
	return cmp.Equal(actual, expected, cmpopts.IgnoreFields(dummy, ignoredFields...))
}
