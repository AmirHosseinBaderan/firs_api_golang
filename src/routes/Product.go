package Server

import (
	"encoding/json"
	Db "first_api/src/database/context"
	Models "first_api/src/database/models"
	"net/http"
	"strconv"
)

type ProductResponse struct {
	Products    []Models.Product
	Pages       int64
	CurrentPage int
}

type Error struct {
	Title   string
	Message string
}

func Getproducts(w http.ResponseWriter, r *http.Request) {
	var products []Models.Product
	var totalRows int64

	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	count, _ := strconv.Atoi(query.Get("count"))

	db := Db.DataBase()

	db.Model(products).Count(&totalRows)
	db.Offset(page).Limit(count).Find(&products)

	var response = ProductResponse{Products: products, Pages: (totalRows / int64(count)), CurrentPage: page}
	json.NewEncoder(w).Encode(response)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product Models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		json.NewEncoder(w).Encode(&Error{Title: "An Exception Throwed When Convert Data", Message: err.Error()})
	} else {
		db := Db.DataBase()
		db.Create(&product)
		json.NewEncoder(w).Encode(product)
	}
}
