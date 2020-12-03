package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

	puppyURL := "http://www.recipepuppy.com/api/"
	base, err := url.Parse(puppyURL)
	if err != nil {
		return recipePuppyResponse, false, "Some wrong happens"
	}

	params := url.Values{}
	params.Add("i", ingredients)
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())

	if err != nil {
		log.Fatal(err)
		return recipePuppyResponse, false, "The RecipePuppy service are not working"
	}

	err = json.NewDecoder(response.Body).Decode(&recipePuppyResponse)
	if err != nil {
		log.Fatal(err)
		return recipePuppyResponse, false, "Some wrong with the response RecipePuppy decode"
	}

	return recipePuppyResponse, true, "every things is fine"

}
