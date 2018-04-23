package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var generator = &IDGenerator{}

func handler(w http.ResponseWriter, request *http.Request) {
	//category := mux.Vars(request)
	value, _ := generator.GetNewId()
	fmt.Fprintf(w, fmt.Sprintf("%d", value))
}

func main() {
	generator.Init()

	//for i := 0; i < 8; i++ {
	//	fmt.Println(generator.GetNewId())
	//}
	router := mux.NewRouter()

	router.HandleFunc("/v1/{category}", handler)
	http.ListenAndServe(":8088", router)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}