package main

import (
	"testing"
)

func TestVar(t *testing.T) {
	got := PrintNumber()
	want := 1

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
