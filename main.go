package main

import (
	"fmt"
	"log"
	"net/http"
	. "sample/routes"
	// "github.com/gorilla/mux"
)

func main(){
	mux := http.NewServeMux()
	userHandle := MakeHTTPHandler()
	mux.Handle("/user", userHandle)
	mux.Handle("/user/", userHandle)
	fmt.Println("Hi here");

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}