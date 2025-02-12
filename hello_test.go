package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Primeiro teste da aninha com go!!"

	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
