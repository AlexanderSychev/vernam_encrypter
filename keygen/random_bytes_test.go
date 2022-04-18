package keygen

import "testing"

func TestRandomBytes(t *testing.T) {
	t.Log("\"RandomBytes\" function tests")

	// Test #1

	t.Log("  Test #1: If length is lower than 1 then function returns empty slice and error object")

	slice01, err01 := RandomBytes(0)
	if len(slice01) > 0 {
		t.Fatalf("  Failed: expected empty slice, got %d byte length slice", len(slice01))
	}
	if err01 == nil {
		t.Fatal("  Expected error object, got \"nil\"")
	}

	t.Log("  Succeed")

	// Test #2

	t.Log("  Test #2: If length is 1 byte or bigger then function returns slice with random bytes")

	for i := 1; i <= 8; i++ {
		t.Logf("    %d bytes length", i)

		slice, err := RandomBytes(i)
		if err != nil {
			t.Fatalf("    Expected \"nil\", got error with message \"%s\"", err.Error())
		}
		if len(slice) != i {
			t.Fatalf("    Expected %d bytes length slice, fot %d bytes length slice", i, len(slice))
		}

		t.Log("    Succeed")
	}

	t.Log("  Succeed")
}
