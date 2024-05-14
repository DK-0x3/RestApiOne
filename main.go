package main

import (
	"RestApiOne/database"
	"RestApiOne/server"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	//fmt.Print(generateRandomString(50))
	r := mux.NewRouter()
	r.HandleFunc("/{api}/main", server.GetMainCategorys).Methods("GET")
	r.HandleFunc("/{api}/mainCat", server.GetMainCategorysAndCategory).Methods("GET")
	r.HandleFunc("/{api}/main/cat", server.GetCategorys).Methods("GET")
	r.HandleFunc("/{api}/main/{id}", server.GetCategoryToMainCategoryID).Methods("GET")
	r.HandleFunc("/{api}/main/cat/{id}", server.GetCategoryID).Methods("GET")
	r.HandleFunc("/{api}/main/cat/product/{id}", server.GetProductToCategoryID).Methods("GET")
	r.HandleFunc("/{api}/api", server.CreateApiKey).Methods("POST")
	r.HandleFunc("/{api}/api", server.GetApiKeyAll).Methods("GET")
	r.HandleFunc("/{api}/api/{key}", server.DeleteApiKey).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}

