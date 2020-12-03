package main

import "testing"

func TestValidateIngredientsParams(t *testing.T) {

	callbackExpected := true
	ingredientsInput := []string{"batata,arroz,tomate"}
	callValidateIngredientsParams(ingredientsInput, callbackExpected, t)
	
	callbackExpected = false
	ingredientsInput = []string{"batata,arroz,tomate,abobora"}
	callValidateIngredientsParams(ingredientsInput, callbackExpected, t)

	callbackExpected = false
	ingredientsInput = []string{""}
	callValidateIngredientsParams(ingredientsInput, callbackExpected, t)

}

func callValidateIngredientsParams(ingredientsInput []string, callbackExpected bool, t *testing.T) {
	callback, _ := getRecipesValidateParams(ingredientsInput, true)
	if callback != callbackExpected {
		t.Errorf("Expected callback to be %v, but got %v", callbackExpected, callback)
	}
}

