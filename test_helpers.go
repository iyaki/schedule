package schedule

import "testing"

func assertEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but expected one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("getted an error but didn't want one")
	}
}
