package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mike")

	got := buffer.String()
	want := "Hello, Mike"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
