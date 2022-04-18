package transform

import "testing"

// testCaseKeyToShortError is test case
type testCaseKeyToShortError struct {
	// Source bytes slice length
	srcLen int
	// Key bytes slice length
	keyLen int
	// Expected "Error()" method result
	expected string
}

func TestKeyToShortError_Error(t *testing.T) {
	t.Log("Tests for \"Error\" method of \"KeyToShortError\" type")

	testCases := []testCaseKeyToShortError{
		{
			srcLen: 3,
			keyLen: 2,
			expected: "Key length is to short: expected at least 3 bytes, got 2 bytes",
		},
		{
			srcLen: 15,
			keyLen: 0,
			expected: "Key length is to short: expected at least 15 bytes, got 0 bytes",
		},
		{
			srcLen: 10,
			keyLen: 4,
			expected: "Key length is to short: expected at least 10 bytes, got 4 bytes",
		},
		{
			srcLen: 256,
			keyLen: 255,
			expected: "Key length is to short: expected at least 256 bytes, got 255 bytes",
		},
	}

	for testIndex, testCase := range testCases {
		t.Logf("  Test #%d", testIndex + 1)
		err := KeyToShortError{
			srcLen: testCase.srcLen,
			keyLen: testCase.keyLen,
		}
		if err.Error() != testCase.expected {
			t.Fatalf("  Failed: expected \"%s\", got \"%s\"", err.Error(), testCase.expected)
		}
		t.Log("  Succeed")
	}
}
