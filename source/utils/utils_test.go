package utils

import "testing"

func TestSplitStringByComma(t *testing.T) {

	ingredientsExpected := []string{"batata", "arroz", "tomate"}
	ingredientsInput := "batata,arroz,tomate"
	callSplitStringByComma(ingredientsInput, ingredientsExpected, t)

	ingredientsExpected = []string{"batata", "arroz", "tomate", "salt", "orion", "potato"}
	ingredientsInput = "batata,arroz,tomate,salt,orion,potato"
	callSplitStringByComma(ingredientsInput, ingredientsExpected, t)

	ingredientsExpected = []string{"batata"}
	ingredientsInput = "batata"
	callSplitStringByComma(ingredientsInput, ingredientsExpected, t)
	
}


func callSplitStringByComma(ingredientsInput string, ingredientsExpected []string, t *testing.T) {
	response := SplitStringByComma(ingredientsInput)

	if response[0] != ingredientsExpected[0] {
		t.Errorf("Expected callback to be %v, but got %v", ingredientsExpected, response)
	}

	if len(response) != len(ingredientsExpected) {
		t.Errorf("Expected callback array length to be %v, but got %v", len(ingredientsExpected), len(response))
	}
}
