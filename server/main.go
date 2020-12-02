package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	ConfigEnvVariables()
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

	resp := GenerateRecipes(ingredients[0])
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}

func getRecipesValidateParams(ingredients []string, ok bool) (bool, string) {
	// validate the ingredients params
	if !ok || len(ingredients[0]) < 1 {
		message := "Url Param 'i' is missing"
		return false, message
	}

	// split the query string by comma
	ingredientsList := SplitStringByComma(ingredients[0])

	// validate the ingredients length
	if len(ingredientsList) > 3 {
		message := "You can not send more than 3 ingredients"
		return false, message
	}

	return true, "every things is fine"

}
