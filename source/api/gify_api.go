package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type GifyResponse struct {
	Data []GifyImagesResponse
}

type GifyImagesResponse struct {
	Images GifyOriginalResponse
}

type GifyOriginalResponse struct {
	Original Gif
}

type Gif struct {
	Url string
}

// CallGifyAPI receive a string word and return a GifyResponse, with a boolean status and a string message
func CallGifyAPI(searchword string) (GifyResponse, bool, string) {

	var gifyResponse GifyResponse

	gifyAPIKey := os.Getenv("GIFY_API_KEY")
	gifyURL := "https://api.giphy.com/v1/gifs/search"

	base, err := url.Parse(gifyURL)
	if err != nil {
		return gifyResponse, false, "Some wrong happens"
	}

	params := url.Values{}
	params.Add("api_key", gifyAPIKey)
	params.Add("q", searchword)
	params.Add("limit", "1")
	params.Add("offset", "0")
	params.Add("rating", "r")
	params.Add("lang", "en")
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		log.Fatal(err)
		return gifyResponse, false, "The Gify service are not working"
	}

	err = json.NewDecoder(response.Body).Decode(&gifyResponse)
	if err != nil {
		log.Fatal(err)
		return gifyResponse, false, "Some wrong with the response RecipePuppy decode"
	}

	return gifyResponse, true, "every things is fine"
}
