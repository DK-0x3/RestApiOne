package database

import (
	"crypto/rand"
	"fmt"
	"math/big"

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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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
//SELECT

// ** MainCategory **
func SelectMainCategoryAll() []MainCategory {
	sqlStr := "select * from maincategory"
	var MainCategorys []MainCategory
	err := db.Select(&MainCategorys, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return MainCategorys
}

// ** Category **
func SelectCategoryAll() []Category {
	sqlStr := "select * from category"
	var Categorys []Category
	err := db.Select(&Categorys, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return Categorys
}
func SelectCategoryId(id int) Category {
	sqlStr := "select * from category where id = $1"
	var result Category
	err := db.Get(&result, sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return result
}

func SelectCategoryToMainCategoryID(id int) []Category {
	sqlStr := "select * from category where idmaincategory = $1"
	var Categorys []Category
	err := db.Select(&Categorys, sqlStr, id)
	if err != nil {
		fmt.Print("Error DB: " + err.Error())
		return nil
	}
	return Categorys
}

// ** Product **
func SelectProductToCategoryID(id int) []Product {
	sqlStr := "select * from product where idcategory = $1"
	var Categorys []Product
	err := db.Select(&Categorys, sqlStr, id)
	if err != nil {
		fmt.Print("Error DB: " + err.Error())
		return nil
	}
	return Categorys
}

// ** API Key **
func SelectApiKeyAll() []ApiKey {
	sqlStr := "select * from apikey"
	var result []ApiKey
	err := db.Select(&result, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return result
}

func CreateApiKey(role string) ApiKey {
	var ApiStr string
	for {
		randomStr := GenerateRandomString(30)
		if GetRoleRequest(randomStr) == "null" {
			ApiStr = randomStr
			break
		}else {
			continue
		}
	}
	
	sqlStr := fmt.Sprintf("insert into apikey(key, role) values ('%s', '%s')", ApiStr, role)
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
	}
	theID, err := ret.LastInsertId() // Self-incrementing ID
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v\n", err)
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return GetApiKey(ApiStr)
}

func DeleteApiKey(key string) bool {
	sqlStr := "delete from apikey where key = $1"
	ret, err := db.Exec(sqlStr, key)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return false
	}
	n, err := ret.RowsAffected() // Number of rows affected by the operation
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return false
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
	return true
}

func GetApiKey(key string) ApiKey {
	sqlStr := "select * from apikey where key = $1"
	var result ApiKey
	err := db.Get(&result, sqlStr, key)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return result
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Определение длины набора символов
	charsetLength := big.NewInt(int64(len(charset)))

	// Создание буфера для хранения случайных байтов
	buf := make([]byte, length)

	// Генерация случайных байтов
	for i := range buf {
		// Генерация случайного индекса для выбора символа из charset
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			panic(err) // Обработка ошибки
		}
		// Преобразование индекса в число в диапазоне от 0 до charsetLength-1
		buf[i] = charset[index.Int64()]
	}

	return string(buf)
}

func GetRoleRequest(API string) string {
	ListApiKeys := SelectApiKeyAll()
	for _, item := range ListApiKeys {
		if item.Key == API {
			return item.Role
		}
	}
	return "null"
}
