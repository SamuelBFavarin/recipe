package main

import "testing"

func TestValidateIngredientsParams(t *testing.T) {
	
	// test if response is good
	callbackExpected := true
	ingredientsInput := []string{"batata,arroz,tomate"}
	okInput := true
	callback, _ := validateIngredientsParams(ingredientsInput, okInput)
	
	if callback != callbackExpected {
		t.Errorf("Expected callback to be %v, but got %v", callbackExpected, callback)
	}

	// test if response is bad
	callbackExpected = false
	ingredientsInput = []string{"batata,arroz,tomate,abobora"}
	okInput = true
	callback, _ = validateIngredientsParams(ingredientsInput, okInput)
	
	if callback != callbackExpected {
		t.Errorf("Expected callback to be %v, but got %v", callbackExpected, callback)
	}

}