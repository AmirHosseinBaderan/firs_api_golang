package Server

import (
	"fmt"
	"net/http"
)

func StartServer(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Cannot Run Server : ", err)
	} else {
		fmt.Println("Server Succesfully Start On : ", addr)
	}
}

func SetRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/getProducts", Getproducts)
	http.HandleFunc("/createProduct", CreateProduct)
}
