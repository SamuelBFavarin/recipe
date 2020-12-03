package api

import ("testing")

func TestPuppyAPI(t *testing.T) {

	responseExpected := true
	ingredientsInput := "apple,salt,banana"
	callPuppyAPI(ingredientsInput, responseExpected, t)

	responseExpected = true
	ingredientsInput = ""
	callPuppyAPI(ingredientsInput, responseExpected, t)

	responseExpected = true
	ingredientsInput = "ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie,ApplePie"
	callPuppyAPI(ingredientsInput, responseExpected, t)
}


func callPuppyAPI(ingredientsInput string, responseExpected bool, t *testing.T) {
	
	_, response, _ := CallPuppyAPI(ingredientsInput)
	if response != responseExpected {
		t.Errorf("Expected response callback to be %v, but got %v", responseExpected, response)
	}

}
