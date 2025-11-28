package main

import "testing"

func TestInfo(t *testing.T) {
	u := User{Name: "Njox", Age: 21}

	got := u.Info()
	want := "Njox"

	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
