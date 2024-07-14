package hasher

import (
	"testing"
)

func TestHashAndCompare(t *testing.T) {
	password := "mysecretpassword"

	hash, err := Hash(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	match, err := Compare(password, hash)
	if err != nil {
		t.Errorf("Error comparing password: %v", err)
	}
	if !match {
		t.Error("Expected passwords to match, but they didn't")
	}

	wrongPassword := "wrongpassword"
	match, err = Compare(wrongPassword, hash)
	if err != nil {
		t.Errorf("Error comparing password: %v", err)
	}
	if match {
		t.Error("Expected passwords not to match, but they did")
	}
}

func TestHashInvalidFormat(t *testing.T) {
	invalidHash := "invalid$format"

	_, err := Compare("password", invalidHash)
	if err == nil {
		t.Error("Expected error for invalid hash format, but got nil")
	}
}
