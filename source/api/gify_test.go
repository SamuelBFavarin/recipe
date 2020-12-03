package api

import ("testing")

func TestGifyAPI(t *testing.T) {

	responseExpected := true
	wordsInput := "Apple Pie"
	callGifyAPI(wordsInput, responseExpected, t)

	responseExpected = true
	wordsInput = "This is a big phrase, but I think I can get a good response"
	callGifyAPI(wordsInput, responseExpected, t)

	responseExpected = true
	wordsInput = "Apple Pie Apple Pie Apple Pie Apple Pie Apple Pie Apple Pie Apple Pie Apple Pie"
	callGifyAPI(wordsInput, responseExpected, t)
}


func callGifyAPI(wordsInput string, responseExpected bool, t *testing.T) {
	
	_, response, _ := CallGifyAPI(wordsInput)
	if response != responseExpected {
		t.Errorf("Expected response callback to be %v, but got %v", responseExpected, response)
	}

}
