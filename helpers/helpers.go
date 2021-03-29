package helpers

import "testing"

func AssertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("didn't expect an error but got one, %s", got)
	}
}

func AssertError(t *testing.T, got error) {
	t.Helper()
	if got == nil {
		t.Fatalf("expected an error but didn't get one")
	}
}

func AssertErrorMessage(t *testing.T, got error, expected string) {
	t.Helper()
	if got.Error() != expected {
		t.Fatalf("received wrong error, expected %s but got %s instead", got.Error(), expected)
	}
}
