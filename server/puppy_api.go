package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RecipePuppyItemResponse struct {
	Title       string
	Href        string
	Ingredients string
	Thumbnail   string
}

type RecipePuppyResponse struct {
	Results []RecipePuppyItemResponse
}

// CallPuppyAPI require a string of ingredients and return a RecipePuppyResponse, with a bool status, and a string message
func CallPuppyAPI(ingredients string) (RecipePuppyResponse, bool, string) {

	var recipePuppyResponse RecipePuppyResponse

	// DO THE REQUEST
	puppyURL := "http://www.recipepuppy.com/api/?i=" + ingredients
	response, err := http.Get(puppyURL)

	if err != nil {
		log.Fatal(err)
		return recipePuppyResponse, false, "The RecipePuppy service are not working"
	}

	// CONVERT THE RESPONSE DATA TO A JSON STRUCT
	err = json.NewDecoder(response.Body).Decode(&recipePuppyResponse)
	if err != nil {
		log.Fatal(err)
		return recipePuppyResponse, false, "Some wrong with the response RecipePuppy decode"
	}

	// RETURN DATA
	return recipePuppyResponse, true, "every things is fine"

}
