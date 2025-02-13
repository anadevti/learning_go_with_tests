package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello("Golang!")
	expected := "Teste, Golang!"

	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
