
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	
	message := printHelloWordMessage()
	fmt.Println(message)
    w.Write([]byte(`{"message":"` + message + `"}`))
}


func printHelloWordMessage() string {
	return "hello world!"
}