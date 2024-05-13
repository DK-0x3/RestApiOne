package database

import (
	"database/sql"
)

type Product struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Price       int 	`db:"price"`
	Discription string  `db:"discription"`
	IdCategory  int     `db:"idcategory"`
}

type Category struct {
	ID             int     `db:"id"`
	Name           string  `db:"name"`
	IdMainCategory int `db:"idmaincategory"`
}

type MainCategory struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Img  sql.NullString `db:"img"`
}

type ApiKey struct {
	Key string `db:"key"`
	Role string `db:"role"`
}