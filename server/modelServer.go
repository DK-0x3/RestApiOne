package server

type ErrorRequest struct {
	ErrorMessage string `json:"errormessage"`
}

type SuccessRequest struct {
	Message string `json:"message"`
}

type MainCategoryJ struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

type MainCategoryAndCategoryJ struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Img      string      `json:"img"`
	Category []CategoryJ `json:"category"`
}

type ProductJ struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Discription string `json:"discription"`
	IdCategory  int    `json:"idcategory"`
}

type CategoryJ struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	IdMainCategory int    `json:"idMainCategory"`
}

type ApiKeyJ struct {
	Key  string `json:"key"`
	Role string `json:"role"`
}