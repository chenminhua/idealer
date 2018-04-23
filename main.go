package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var service *IDService

func handler(w http.ResponseWriter, request *http.Request) {
	category := mux.Vars(request)["category"]
	value, _ := service.GetNewId(category)
	fmt.Fprintf(w, fmt.Sprintf("%d", value))
}

func main() {
	InitDB()
	service = CreateIdService()
	router := mux.NewRouter()
	router.HandleFunc("/v1/{category}", handler)
	http.ListenAndServe(":8088", router)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}