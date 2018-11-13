package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHomePage(w http.ResponseWriter, r*http.Request){
	fmt.Fprintf(w, "Wlcome to my home page")
	fmt.Println("Endpoint Hit:myHomePage")

}
func handleHttpRequest(){
	http.HandleFunc("/", myHomePage)
	log.Fatal(http.ListenAndServe( ":8080", nil))
}

func main(){
	handleHttpRequest()

}