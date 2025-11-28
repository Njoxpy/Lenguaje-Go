package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(12, 13)
	want := 25

	if got != want {
		t.Errorf("Got: %d, Want: %d", got, want) // use %d for integers
	}
}
