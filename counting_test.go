package main

import "testing"

func TestCount(t *testing.T) {
	got := Count("hello")
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
