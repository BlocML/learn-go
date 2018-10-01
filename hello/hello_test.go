package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tinashe")
	want := "hello, Tinashe"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
