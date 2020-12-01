package main

import "testing"

func TestPrintHelloWordMessage(t *testing.T) {
	outputExpected := "hello world!"

	actual := printHelloWordMessage()

	
	if actual != outputExpected {
		t.Errorf("Expected output to be %s, but got %s", outputExpected, actual)
	}
}