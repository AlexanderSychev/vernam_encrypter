package transform

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	t.Log("\"Decrypt\" function tests")

	// Test #1

	t.Log("  Test #1: if key length is lower than source length then empty slice and error will return")

	src01, key01 := []byte{5, 6, 33}, []byte{10, 4}
	result01, err01 := Decrypt(src01, key01)

	if len(result01) > 0 {
		t.Fatalf("  Failed: expected empty slice, got %d length slice", len(result01))
	}

	if err01 == nil {
		t.Fatal("  Failed: expected error object, got \"nil\"")
	}

	t.Log("  Succeed")

	// Test #2

	t.Log("  Test #2: if key length equals source length then slice with source length will be returned")

	src02, key02 := []byte{5, 6, 33}, []byte{10, 4, 22}
	expectedResult01 := []byte{15, 2, 55}
	result02, err02 := Decrypt(src02, key02)

	if err02 != nil {
		t.Fatalf("  Failed: expected \"nil\" got error object with message \"%s\"", err02.Error())
	}

	if len(result02) != len(src02) {
		t.Fatalf("  Failed: expected %d length slice, got %d length slice", len(src02), len(result02))
	}

	if !reflect.DeepEqual(result02, expectedResult01) {
		t.Fatalf("  Failed: expected %v slice, got %v", expectedResult01, result02)
	}

	t.Log("  Succeed")

	// Test #3

	t.Log("  Test #3: if key length bigger than source length then slice with source length will be returned")

	src03, key03 := []byte{10, 37, 118}, []byte{25, 34, 125, 200}
	expectedResult02 := []byte{19, 7, 11}
	result03, err03 := Decrypt(src03, key03)

	if err03 != nil {
		t.Fatalf("  Failed: expected \"nil\" got error object with message \"%s\"", err03.Error())
	}

	if len(result03) != len(src03) {
		t.Fatalf("  Failed: expected %d length slice, got %d length slice", len(src03), len(result03))
	}

	if !reflect.DeepEqual(result03, expectedResult02) {
		t.Fatalf("  Failed: expected %v slice, got %v", expectedResult02, result03)
	}

	t.Log("  Succeed")
}
