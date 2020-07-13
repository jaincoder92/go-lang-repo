package main

import (
	"fmt"
	"github.com/go-lang-repo/ORM/handlers"
	"github.com/gorilla/mux"
)

func handleRequests(){
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/", HandleRootAPI)
}

func main(){
	fmt.Println("GORM Project client has been found.")
	handleRequests()

}
