package main

import (
	"fmt"
	"log"
	"net/http"
	entity "github.com/SamuelBFavarin/recipe/source/entity"
	utils "github.com/SamuelBFavarin/recipe/source/utils"
	config "github.com/SamuelBFavarin/recipe/source/config"

)

func main() {

	config.ConfigEnvVariables()
	endpoints()
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func endpoints() {
	http.HandleFunc("/recipes", getRecipes)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	ingredients, ok := r.URL.Query()["i"]

	validateStatus, message := getRecipesValidateParams(ingredients, ok)

	if validateStatus == false {
		w.Write([]byte(`{"message":"` + message + `"}`))
		return
	}

	resp := entity.GenerateRecipes(ingredients[0])
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}

func getRecipesValidateParams(ingredients []string, ok bool) (bool, string) {
	// validate the ingredients params
	if !ok || len(ingredients[0]) < 1 {
		message := "Url Param 'i' is missing"
		return false, message
	}

	// validate the ingredients length
	ingredientsList := utils.SplitStringByComma(ingredients[0])
	if len(ingredientsList) > 3 {
		message := "You can not send more than 3 ingredients"
		return false, message
	}

	return true, "every things is fine"

}
