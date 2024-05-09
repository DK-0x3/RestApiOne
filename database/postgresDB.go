package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "128775"
    dbname   = "GoShop"
)

func InitDB() (err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
        					host, port, user, password, dbname)
	
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func SelectMainCategoryAll() []MainCategory{
	sqlStr := "select * from maincategory"
	var MainCategorys []MainCategory
	err := db.Select(&MainCategorys, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return MainCategorys
	// for _, value := range MainCategorys {
	// 	fmt.Println(fmt.Sprintf("ID: %d Name: %s img: %s", value.ID, value.Name, value.Img.String))
	// }
}