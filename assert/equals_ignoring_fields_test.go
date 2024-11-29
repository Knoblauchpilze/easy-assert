package assert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EqualsIgnoringFields_Int(t *testing.T) {
	compareInts := func() {
		EqualsIgnoringFields(0, 0)
	}

	assert.Panics(t, compareInts)
}

func Test_EqualsIgnoringFields_Time(t *testing.T) {
	type testCase struct {
		t1       time.Time
		t2       time.Time
		expected bool
	}

	berlinTime, err := time.LoadLocation("Europe/Berlin")
	assert.Nil(t, err)
	parisTime, err := time.LoadLocation("Europe/Paris")
	assert.Nil(t, err)
	londonTime, err := time.LoadLocation("Europe/London")
	assert.Nil(t, err)

	testCases := []testCase{
		{
			t1:       time.Date(2024, 11, 29, 15, 13, 43, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 15, 13, 43, 0, time.UTC),
			expected: true,
		},
		{
			t1:       time.Date(2024, 11, 29, 15, 13, 43, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 15, 13, 44, 0, time.UTC),
			expected: false,
		},
		{
			t1:       time.Date(2024, 11, 29, 15, 13, 43, 0, berlinTime),
			t2:       time.Date(2024, 11, 29, 15, 13, 43, 0, parisTime),
			expected: true,
		},
		{
			t1:       time.Date(2024, 11, 29, 15, 13, 43, 0, berlinTime),
			t2:       time.Date(2024, 11, 29, 15, 13, 43, 0, londonTime),
			expected: false,
		},
		{
			t1:       time.Date(2024, 11, 29, 15, 13, 43, 0, berlinTime),
			t2:       time.Date(2024, 11, 29, 14, 13, 43, 0, londonTime),
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.t1, testCase.t2),
			)
		})
	}
}

func Test_EqualsIgnoringFields_Struct(t *testing.T) {
	type dummyStruct struct {
		A    int
		Name string
	}

	type testCase struct {
		lhs      dummyStruct
		rhs      dummyStruct
		expected bool
	}

	testCases := []testCase{
		{
			lhs:      dummyStruct{},
			rhs:      dummyStruct{},
			expected: true,
		},
		{
			lhs:      dummyStruct{A: 1},
			rhs:      dummyStruct{},
			expected: false,
		},
		{
			lhs:      dummyStruct{},
			rhs:      dummyStruct{A: 1},
			expected: false,
		},
		{
			lhs:      dummyStruct{Name: "lhs"},
			rhs:      dummyStruct{},
			expected: false,
		},
		{
			lhs:      dummyStruct{},
			rhs:      dummyStruct{Name: "rhs"},
			expected: false,
		},
		{
			lhs:      dummyStruct{A: 123},
			rhs:      dummyStruct{A: 123},
			expected: true,
		},
		{
			lhs:      dummyStruct{A: 1},
			rhs:      dummyStruct{A: -987198},
			expected: false,
		},
		{
			lhs:      dummyStruct{Name: "equal"},
			rhs:      dummyStruct{Name: "equal"},
			expected: true,
		},
		{
			lhs:      dummyStruct{Name: "lhslhs"},
			rhs:      dummyStruct{Name: "rhsrhs"},
			expected: false,
		},
		{
			lhs:      dummyStruct{A: 1, Name: "thoseAreEqual"},
			rhs:      dummyStruct{A: 2, Name: "thoseAreEqual"},
			expected: false,
		},
		{
			lhs:      dummyStruct{A: -19, Name: "value1"},
			rhs:      dummyStruct{A: -19, Name: "value2"},
			expected: false,
		},
		{
			lhs:      dummyStruct{A: 654, Name: "identical"},
			rhs:      dummyStruct{A: 654, Name: "identical"},
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs),
			)
		})
	}
}

func Test_EqualsIgnoringFields_StructWithTime(t *testing.T) {
	type dummyStruct struct {
		Name      string
		CreatedAt time.Time
	}

	type testCase struct {
		lhs      dummyStruct
		rhs      dummyStruct
		expected bool
	}

	berlinTime, err := time.LoadLocation("Europe/Berlin")
	assert.Nil(t, err)
	londonTime, err := time.LoadLocation("Europe/London")
	assert.Nil(t, err)

	testCases := []testCase{
		{
			lhs: dummyStruct{
				Name:      "name1",
				CreatedAt: time.Date(2024, 11, 29, 15, 28, 06, 0, time.UTC),
			},
			rhs: dummyStruct{
				Name:      "name2",
				CreatedAt: time.Date(2024, 11, 29, 15, 28, 06, 0, time.UTC),
			},
			expected: false,
		},
		{
			lhs: dummyStruct{
				Name:      "name",
				CreatedAt: time.Date(2024, 11, 29, 15, 28, 31, 0, time.UTC),
			},
			rhs: dummyStruct{
				Name:      "name",
				CreatedAt: time.Date(2024, 11, 29, 15, 28, 06, 0, time.UTC),
			},
			expected: false,
		},
		{
			lhs: dummyStruct{
				Name:      "name",
				CreatedAt: time.Date(2024, 11, 29, 15, 28, 31, 0, berlinTime),
			},
			rhs: dummyStruct{
				Name:      "name",
				CreatedAt: time.Date(2024, 11, 29, 14, 28, 31, 0, londonTime),
			},
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs),
			)
		})
	}
}

func Test_EqualsIgnoringFields_NestedStruct(t *testing.T) {
	type nestedStruct struct {
		A     int
		Value string
	}

	type dummyStruct struct {
		Name   string
		Nested nestedStruct
	}

	type testCase struct {
		lhs      dummyStruct
		rhs      dummyStruct
		expected bool
	}

	testCases := []testCase{
		{
			lhs:      dummyStruct{},
			rhs:      dummyStruct{},
			expected: true,
		},
		{
			lhs: dummyStruct{
				Name: "lhs",
			},
			rhs:      dummyStruct{},
			expected: false,
		},
		{
			lhs: dummyStruct{},
			rhs: dummyStruct{
				Name: "rhs",
			},
			expected: false,
		},
		{
			lhs: dummyStruct{
				Nested: nestedStruct{},
			},
			rhs:      dummyStruct{},
			expected: true,
		},
		{
			lhs: dummyStruct{
				Nested: nestedStruct{
					A: 1,
				},
			},
			rhs:      dummyStruct{},
			expected: false,
		},
		{
			lhs: dummyStruct{
				Nested: nestedStruct{
					A: 1,
				},
			},
			rhs: dummyStruct{
				Nested: nestedStruct{
					A: 1,
				},
			},
			expected: true,
		},
		{
			lhs: dummyStruct{
				Nested: nestedStruct{
					A:     1,
					Value: "nested1",
				},
			},
			rhs: dummyStruct{
				Nested: nestedStruct{
					A:     1,
					Value: "nested2",
				},
			},
			expected: false,
		},
		{
			lhs: dummyStruct{
				Nested: nestedStruct{
					A:     2,
					Value: "nested",
				},
			},
			rhs: dummyStruct{
				Nested: nestedStruct{
					A:     1,
					Value: "nested",
				},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs),
			)
		})
	}
}

func Test_EqualsIgnoringFields_IgnoringOneField(t *testing.T) {
	type dummyStruct struct {
		Name string
		A    float32
	}

	type testCase struct {
		lhs          dummyStruct
		rhs          dummyStruct
		ignoredField string
		expected     bool
	}

	testCases := []testCase{
		{
			lhs:          dummyStruct{},
			rhs:          dummyStruct{},
			ignoredField: "A",
			expected:     true,
		},
		{
			lhs:          dummyStruct{A: 1},
			rhs:          dummyStruct{A: 2},
			ignoredField: "A",
			expected:     true,
		},
		{
			lhs:          dummyStruct{A: 1},
			rhs:          dummyStruct{A: 2},
			ignoredField: "Name",
			expected:     false,
		},
		{
			lhs:          dummyStruct{A: 1},
			rhs:          dummyStruct{A: 2, Name: "rhs"},
			ignoredField: "Name",
			expected:     false,
		},
		{
			lhs:          dummyStruct{A: 39, Name: "rhs"},
			rhs:          dummyStruct{A: 39},
			ignoredField: "Name",
			expected:     true,
		},
		{
			lhs:          dummyStruct{A: 39, Name: "name"},
			rhs:          dummyStruct{A: 39, Name: "name"},
			ignoredField: "Name",
			expected:     true,
		},
		{
			lhs:          dummyStruct{A: 39, Name: "name"},
			rhs:          dummyStruct{A: 39, Name: "name"},
			ignoredField: "A",
			expected:     true,
		},
		{
			lhs:          dummyStruct{A: 38, Name: "lhs"},
			rhs:          dummyStruct{A: 39, Name: "rhs"},
			ignoredField: "A",
			expected:     false,
		},
		{
			lhs:          dummyStruct{A: 29, Name: "value"},
			rhs:          dummyStruct{A: 41, Name: "value"},
			ignoredField: "A",
			expected:     true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs, testCase.ignoredField),
			)
		})
	}
}

func Test_EqualsIgnoringFields_IgnoringMultipleFields(t *testing.T) {
	type dummyStruct struct {
		Name      string
		A         float32
		CreatedAt time.Time
	}

	type testCase struct {
		lhs           dummyStruct
		rhs           dummyStruct
		ignoredFields []string
		expected      bool
	}

	berlinTime, err := time.LoadLocation("Europe/Berlin")
	assert.Nil(t, err)
	londonTime, err := time.LoadLocation("Europe/London")
	assert.Nil(t, err)

	testCases := []testCase{
		{
			lhs:           dummyStruct{},
			rhs:           dummyStruct{},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, time.UTC),
			},
			rhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, time.UTC),
			},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, berlinTime),
			},
			rhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, londonTime),
			},
			ignoredFields: []string{"A", "Name"},
			expected:      false,
		},
		{
			lhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, berlinTime),
			},
			rhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, londonTime),
			},
			ignoredFields: []string{"Name", "CreatedAt"},
			expected:      true,
		},
		{
			lhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 15, 40, 10, 0, berlinTime),
			},
			rhs: dummyStruct{
				CreatedAt: time.Date(2024, 11, 29, 14, 40, 10, 0, londonTime),
			},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs, testCase.ignoredFields...),
			)
		})
	}
}

func Test_EqualsIgnoringFields_IgnoringAllFields(t *testing.T) {
	type dummyStruct struct {
		Name string
		A    float32
	}

	type testCase struct {
		lhs           dummyStruct
		rhs           dummyStruct
		ignoredFields []string
		expected      bool
	}

	testCases := []testCase{
		{
			lhs:           dummyStruct{},
			rhs:           dummyStruct{},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs:           dummyStruct{A: 1},
			rhs:           dummyStruct{A: 2},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs:           dummyStruct{A: 1},
			rhs:           dummyStruct{A: 2, Name: "rhs"},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs:           dummyStruct{A: 39, Name: "rhs"},
			rhs:           dummyStruct{A: 39},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs:           dummyStruct{A: 39, Name: "name"},
			rhs:           dummyStruct{A: 39, Name: "name"},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
		{
			lhs:           dummyStruct{A: 456, Name: "lhs"},
			rhs:           dummyStruct{A: 369, Name: "rhs"},
			ignoredFields: []string{"A", "Name"},
			expected:      true,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				EqualsIgnoringFields(testCase.lhs, testCase.rhs, testCase.ignoredFields...),
			)
		})
	}
}
