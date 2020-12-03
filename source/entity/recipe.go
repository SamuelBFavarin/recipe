package entity

import (
	"encoding/json"
	"sort"
	api "github.com/SamuelBFavarin/recipe/source/api"
	utils "github.com/SamuelBFavarin/recipe/source/utils"

)

type RecipeResponse struct {
	Keywords []string
	Recipes  []Recipe
}

type Recipe struct {
	Title       string
	Ingredients []string
	Link        string
	Gif         string
}

// GenerateRecipes used when /recipes endpoit is called. Require a "i" querystring
func GenerateRecipes(ingredients string) []byte {
	response, _, _ := api.CallPuppyAPI(ingredients)
	result := buildResponseAPI(ingredients, response)
	resp, _ := json.Marshal(result)
	return resp
}

func buildResponseAPI(keywords string, puppyResponse api.RecipePuppyResponse) RecipeResponse {

	var recipes []Recipe

	for _, recipe := range puppyResponse.Results {

		gif, _, _ := api.CallGifyAPI(recipe.Title)

		ingredients := utils.SplitStringByComma(recipe.Ingredients)
		sort.Strings(ingredients)

		recipeResponse := Recipe{
			Title:       recipe.Title,
			Ingredients: ingredients,
			Link:        recipe.Href,
			Gif:         gif.Data[0].Images.Original.Url}

		recipes = append(recipes, recipeResponse)
	}

	return RecipeResponse{
		Keywords: utils.SplitStringByComma(keywords),
		Recipes:  recipes}

}
