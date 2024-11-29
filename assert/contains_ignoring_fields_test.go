package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainsIgnoringFields_NotContained(t *testing.T) {
	type dummyStruct struct {
		A    int
		Name string
	}

	actual := []dummyStruct{
		{
			A:    1,
			Name: "value1",
		},
		{
			A:    2,
			Name: "somethingElse",
		},
	}

	expected := dummyStruct{
		A:    0,
		Name: "value1",
	}

	assert.False(t, ContainsIgnoringFields(actual, expected))
}

func Test_ContainsIgnoringFields_Contained(t *testing.T) {
	type dummyStruct struct {
		A    int
		Name string
	}

	actual := []dummyStruct{
		{
			A:    1,
			Name: "value1",
		},
		{
			A:    2,
			Name: "somethingElse",
		},
	}

	expected := dummyStruct{
		A:    1,
		Name: "value1",
	}

	assert.True(t, ContainsIgnoringFields(actual, expected))
}

func Test_ContainsIgnoringFields_ContainedWithIgnore(t *testing.T) {
	type dummyStruct struct {
		A    int
		Name string
	}

	actual := []dummyStruct{
		{
			A:    1,
			Name: "value1",
		},
		{
			A:    2,
			Name: "somethingElse",
		},
	}

	expected := dummyStruct{
		A:    0,
		Name: "value1",
	}

	assert.True(t, ContainsIgnoringFields(actual, expected, "A"))
}
