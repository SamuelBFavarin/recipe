
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"		
	"sort"
	"strings"
)

// RECIPE STRUCT
type RecipeResponse struct {
	Keywords	[]string
	Recipes 	[]Recipe
}

type Recipe struct {
	Title 		string
	Ingredients []string
	Link 		string
	Gif 		string
}

// PUPPY RESPONSE STRUCT
type RecipePuppyItemResponse struct {
	Title		string
	Href		string
	Ingredients string
	Thumbnail	string
}

type RecipePuppyResponse struct {
	Results		[]RecipePuppyItemResponse
} 

// GIFY RESPONSE STRUCT
type GifyResponse struct {
	Data		[]GifyImagesResponse
}

type GifyImagesResponse struct {
	Images		GifyOriginalResponse		
}

type GifyOriginalResponse struct {
	Original	Gif		
}

type Gif struct {
	Url			string
}


func main() {
	// endpoints
	http.HandleFunc("/recipes", recipes)

    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func recipes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	ingredients, ok := r.URL.Query()["i"]
	callback, message := validateIngredientsParams(ingredients, ok)
	
	if callback == false {
		w.Write([]byte(`{"message":"` + message + `"}`))
		return 
	} else {
		response, _, _ := callPuppyAPI(ingredients[0])
		result := buildResponseAPI(ingredients[0], response, "this will be a gify data")
		resp, _ := json.Marshal(result)

		w.Header().Set("Content-Type", "application/json")
    	w.Write(resp)
	}

}

func callPuppyAPI(ingredients string) (RecipePuppyResponse, bool, string ) {
	
	var recipePuppyResponse RecipePuppyResponse

	// DO THE REQUEST
	puppy_url := "http://www.recipepuppy.com/api/?i=" + ingredients
	response, err := http.Get(puppy_url)

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

func callGifyAPI(search_word string) (GifyResponse, bool, string) {
	
	var gify_response GifyResponse

	gify_api_key := "Vkq266MEDGzUYkMrufXHizf5sHxUPPvd"
	gify_url := "https://api.giphy.com/v1/gifs/search"
	
	base, err := url.Parse(gify_url)
    if err != nil {
        return gify_response, false, "Some wrong happens"
    }    
	
	params := url.Values{}		
	params.Add("api_key", gify_api_key)
	params.Add("q", search_word)
	params.Add("limit", "1")
	params.Add("offset", "0")
	params.Add("rating", "r")
	params.Add("lang", "en")
    base.RawQuery = params.Encode() 


	response, err := http.Get(base.String())
	if err != nil {
		log.Fatal(err)
		return gify_response, false, "The Gify service are not working"
	}
	
	err = json.NewDecoder(response.Body).Decode(&gify_response)
	if err != nil {
		log.Fatal(err)
		return gify_response, false, "Some wrong with the response RecipePuppy decode"
	}	

	return gify_response, true, "every things is fine"
}

func buildResponseAPI(keywords string, puppy_response RecipePuppyResponse, gify_response string ) RecipeResponse {

	var recipes []Recipe

	for _, recipe := range puppy_response.Results {
		
		gif, _, _ := callGifyAPI(recipe.Title)

		ingredients := splitStringByComma(recipe.Ingredients)
		sort.Strings(ingredients)

		recipe_response := Recipe{
			Title : recipe.Title,
			Ingredients : ingredients,
			Link : recipe.Href,
			Gif : gif.Data[0].Images.Original.Url}

		recipes = append(recipes, recipe_response)				
	}

	return RecipeResponse{
		Keywords : splitStringByComma(keywords),
		Recipes: recipes}

}

func validateIngredientsParams(ingredients [] string, ok bool) (bool, string) {
	// validate the ingredients params
	if !ok || len(ingredients[0]) < 1 {
		message := "Url Param 'i' is missing"
        return false, message
	}

	// split the query string by comma
	ingredients_list := splitStringByComma(ingredients[0])

	// validate the ingredients length
	if len(ingredients_list) > 3 {
		message := "You can not send more than 3 ingredients"
        return false, message
	}

	return true, "every things is fine"

}

func splitStringByComma(str string) []string {
	cleaned := strings.Replace(str, ",", " ", -1)
	return strings.Fields(cleaned)
}