package main

import (
	"testing"
)

const f string = "./test.txt"

// counts characters including multibyte characters
func TestMflag(t *testing.T) {
	got := countChars(f)
	want := "339292"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestCFlag(t *testing.T) {
	got := countBytes(f)
	want := "342190"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestLFlag(t *testing.T) {
	got := countLines(f)
	want := "7145"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestWFlag(t *testing.T) {
	got := countWords(f)
	want := "58164"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
