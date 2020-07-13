package handlers

import (
	"fmt"
	"net/http"
)

func HandleRootAPI(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello Root API")
}
