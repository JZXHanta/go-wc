package main

import (
	"testing"
)

const testFile string = "./test.txt"

// counts characters including multibyte characters
func TestMflag(t *testing.T) {
	got := countChars(testFile, nil)
	want := "339292"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestCflag(t *testing.T) {
	got := countBytes(testFile, nil)
	want := "342190"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestLflag(t *testing.T) {
	got := countLines(testFile, nil)
	want := "7145"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestWflag(t *testing.T) {
	got := countWords(testFile, nil)
	want := "58164"

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
