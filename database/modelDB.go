package database

import (
	"database/sql"
)

type product struct {
	id          int     `db:"id"`
	name        string  `db:"name"`
	price       float64 `db:"price"`
	discription string  `db:"discription"`
	idCategory  int     `db:"idCategory"`
}

type category struct {
	id             int     `db:"id"`
	name           string  `db:"name"`
	idMainCategory float64 `db:"idMainCategory"`
}

type MainCategory struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Img  sql.NullString `db:"img"`
}