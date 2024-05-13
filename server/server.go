package server

import (
	"RestApiOne/database"
	"encoding/json"
	"fmt"
	"strconv"

	//"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	//"strconv"
)

// -- MainCategory --

func GetMainCategorys(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "user" || database.GetRoleRequest(params["api"]) == "admin") {
		var MainCategorys []MainCategoryJ = mainCategoryToMainCategoryJ(database.SelectMainCategoryAll())
    	json.NewEncoder(w).Encode(MainCategorys)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}

// -- Category --
func GetCategorys(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "user" || database.GetRoleRequest(params["api"]) == "admin") {
		var Categorys []CategoryJ = CategoryToCategoryJ(database.SelectCategoryAll())
    	json.NewEncoder(w).Encode(Categorys)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}

func GetCategoryToMainCategoryID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "user" || database.GetRoleRequest(params["api"]) == "admin") {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Print("Error: " + err.Error())
		}
    	var Category []CategoryJ = CategoryToCategoryJ(database.SelectCategoryToMainCategoryID(id))
    	json.NewEncoder(w).Encode(Category)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
	
}

func GetCategoryID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "user" || database.GetRoleRequest(params["api"]) == "admin") {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Print("Error: " + err.Error())
		}
		x := []database.Category{database.SelectCategoryId(id)}
    	var Category []CategoryJ = CategoryToCategoryJ(x)
    	json.NewEncoder(w).Encode(Category)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
	
}

// -- Product --
func GetProductToCategoryID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "user" || database.GetRoleRequest(params["api"]) == "admin") {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Print("Error: " + err.Error())
		}
    	var Category []ProductJ = ProductToProductJ(database.SelectProductToCategoryID(id))
    	json.NewEncoder(w).Encode(Category)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}





func mainCategoryToMainCategoryJ(mainCategory []database.MainCategory) []MainCategoryJ {
	var result []MainCategoryJ
	
	for i := 0; i < len(mainCategory); i++ {
		result = append(result, MainCategoryJ{
			ID:   mainCategory[i].ID,
			Name: mainCategory[i].Name,
			Img:  mainCategory[i].Img.String,
		})
	}
	return result
}
func CategoryToCategoryJ(Category []database.Category) []CategoryJ {
	var result []CategoryJ
	
	for i := 0; i < len(Category); i++ {
		result = append(result, CategoryJ{
			ID: Category[i].ID,
			Name: Category[i].Name,
			IdMainCategory: Category[i].IdMainCategory,
		})
	}
	return result
}
func ProductToProductJ(product []database.Product) []ProductJ {
	var result []ProductJ
	
	for i := 0; i < len(product); i++ {
		result = append(result, ProductJ{
			ID: product[i].ID,
			Name: product[i].Name,
			IdCategory: product[i].IdCategory,
			Price: product[i].Price,
			Discription: product[i].Discription,
		})
	}
	return result
}

// -- API Key --

func CreateApiKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "admin") {
		var key ApiKeyJ
		_ = json.NewDecoder(r.Body).Decode(&key)
		key.Key = database.CreateApiKey(key.Role).Key
		json.NewEncoder(w).Encode(key)
	} else if database.GetRoleRequest(params["api"]) == "user" {
		message := ErrorRequest{ErrorMessage: "Нет доступа для *User* | There is no access for the User"}
		json.NewEncoder(w).Encode(message)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}

func GetApiKeyAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "admin") {
		
    	var Category []ApiKeyJ = ApiKeyToApiKeyJ(database.SelectApiKeyAll())
    	json.NewEncoder(w).Encode(Category)

	} else if database.GetRoleRequest(params["api"]) == "user" {
		message := ErrorRequest{ErrorMessage: "Нет доступа для *User* | There is no access for the User"}
		json.NewEncoder(w).Encode(message)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteApiKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (database.GetRoleRequest(params["api"]) == "admin") {

    	if database.DeleteApiKey(params["key"]) {
			message := SuccessRequest{Message: "Успешно удалено | Successfully deleted"}
			json.NewEncoder(w).Encode(message)
		}

	} else if database.GetRoleRequest(params["api"]) == "user" {
		message := ErrorRequest{ErrorMessage: "Нет доступа для *User* | There is no access for the User"}
		json.NewEncoder(w).Encode(message)
	} else {
		message := ErrorRequest{ErrorMessage: "Нет доступа, необходим API Ключ | There is no access, an API Key is required"}
		json.NewEncoder(w).Encode(message)
	}
}

func ApiKeyToApiKeyJ(key []database.ApiKey) []ApiKeyJ {
	var result []ApiKeyJ
	
	for i := 0; i < len(key); i++ {
		result = append(result, ApiKeyJ{
			Key: key[i].Key,
			Role: key[i].Role,
		})
	}
	return result
}