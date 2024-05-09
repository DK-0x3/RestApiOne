package server

import (
	"RestApiOne/database"
	"encoding/json"

	//"math/rand"
	"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
)



func GetMainCategorys(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var MainCategorys []MainCategoryJ = mainCategoryToMainCategoryJ(database.SelectMainCategoryAll())
    json.NewEncoder(w).Encode(MainCategorys)
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