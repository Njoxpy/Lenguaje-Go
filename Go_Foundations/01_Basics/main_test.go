package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello world"

	if got != want {
		t.Errorf("HelloWorld() = %q; want %q", got, want)
	}

}
