
package main

import (
    "fmt"
	"log"
	"strings"
    "net/http"
)

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
		w.Header().Set("Content-Type", "application/json")
		message := "hey this is the ingredients: " + ingredients[0]
    	w.Write([]byte(`{"message":"` + message + `"}`))
	}

}

func validateIngredientsParams(ingredients [] string, ok bool) (bool, string) {
	// validate the ingredients params
	if !ok || len(ingredients[0]) < 1 {
		message := "Url Param 'i' is missing"
        return false, message
	}

	// split the query string by comma
	cleaned := strings.Replace(ingredients[0], ",", " ", -1)
	ingredients_list := strings.Fields(cleaned)

	// validate the ingredients length
	if len(ingredients_list) > 3 {
		message := "You can not send more than 3 ingredients"
        return false, message
	}

	return true, "every things is fine"

}