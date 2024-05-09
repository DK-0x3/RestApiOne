package main

import (
	"RestApiOne/database"
	"RestApiOne/server"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()
	fmt.Print(database.SelectMainCategoryAll())
	
	r.HandleFunc("/main", server.GetMainCategorys).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}