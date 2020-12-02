package main

import (
	"encoding/json"
	"sort"
)

// RECIPE STRUCT
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
	response, _, _ := CallPuppyAPI(ingredients)
	result := buildResponseAPI(ingredients, response)
	resp, _ := json.Marshal(result)
	return resp

}

func buildResponseAPI(keywords string, puppyResponse RecipePuppyResponse) RecipeResponse {

	var recipes []Recipe

	for _, recipe := range puppyResponse.Results {

		gif, _, _ := CallGifyAPI(recipe.Title)

		ingredients := SplitStringByComma(recipe.Ingredients)
		sort.Strings(ingredients)

		recipeResponse := Recipe{
			Title:       recipe.Title,
			Ingredients: ingredients,
			Link:        recipe.Href,
			Gif:         gif.Data[0].Images.Original.Url}

		recipes = append(recipes, recipeResponse)
	}

	return RecipeResponse{
		Keywords: SplitStringByComma(keywords),
		Recipes:  recipes}

}
