package keygen

import "testing"

// testCaseSliceToShortError is
type testCaseSliceToShortError struct {
	// Received length of slice to generate
	length int
	// Expected result of "Error" method
	expected string
}

func TestSliceToShortError_Error(t *testing.T) {
	t.Log("Tests for \"Error\" method of \"SliceToShortError\" type")

	testCases := []testCaseSliceToShortError{
		{
			length: -2,
			expected: "Cannot generate bytes slice with length -2. Length must be at least 1 byte",
		},
		{
			length: -1,
			expected: "Cannot generate bytes slice with length -1. Length must be at least 1 byte",
		},
		{
			length: 0,
			expected: "Cannot generate bytes slice with length 0. Length must be at least 1 byte",
		},
	}

	for testIndex, testCase := range testCases {
		err := SliceToShortError{
			length: testCase.length,
		}

		t.Logf("  Test #%d", testIndex + 1)
		if err.Error() != testCase.expected {
			t.Fatalf("  Failed: expected \"%s\", got \"%s\"", testCase.expected, err.Error())
		}
		t.Logf("  Succeed")
	}
}
