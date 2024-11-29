package assert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_AreTimeCloserThan(t *testing.T) {
	type testCase struct {
		t1       time.Time
		t2       time.Time
		distance time.Duration
		expected bool
	}

	testCases := []testCase{
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 0,
			expected: true,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 1 * time.Second,
			expected: true,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 56, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 1 * time.Second,
			expected: true,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 57, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 1 * time.Second,
			expected: false,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 57, 0, time.UTC),
			distance: 1 * time.Second,
			expected: false,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 56, 1, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 1 * time.Second,
			expected: false,
		},
		{
			t1:       time.Date(2024, 11, 29, 18, 12, 56, 1, time.UTC),
			t2:       time.Date(2024, 11, 29, 18, 12, 55, 0, time.UTC),
			distance: 1 * time.Second,
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(
				t,
				testCase.expected,
				AreTimeCloserThan(testCase.t1, testCase.t2, testCase.distance),
			)
		})
	}
}
