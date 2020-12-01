
package main

import (
    "fmt"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
	Results	[]RecipePuppyItemResponse
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

	data, _ := ioutil.ReadAll( response.Body )
	response.Body.Close()

	// CONVERT THE RESPONSE DATA TO A JSON STRUCT
	err = json.Unmarshal(data, &recipePuppyResponse)
	if err != nil {
		log.Fatal(err)
		return recipePuppyResponse, false, "Some wrong with the response RecipePuppy decode"
	}

	// RETURN DATA
	return recipePuppyResponse, true, "every things is fine"

}

func buildResponseAPI(keywords string, puppy_response RecipePuppyResponse, gify_response string ) RecipeResponse {

	var recipes []Recipe

	for _, recipe := range puppy_response.Results {
		
		recipe_response := Recipe{
			Title : recipe.Title,
			Ingredients : splitStringByComma(recipe.Ingredients),
			Link : recipe.Href,
			Gif : recipe.Thumbnail}

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