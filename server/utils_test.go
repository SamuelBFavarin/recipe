package main

import "testing"

func TestSplitStringByComma(t *testing.T) {

	ingredientsExpected := []string{"batata", "arroz", "tomate"}
	ingredientsInput := "batata,arroz,tomate"

	response := SplitStringByComma(ingredientsInput)

	if response[0] != ingredientsExpected[0] {
		t.Errorf("Expected callback to be %v, but got %v", ingredientsExpected, response)
	}
}
