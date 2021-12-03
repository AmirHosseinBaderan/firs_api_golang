package main

import (
	C "first_api/src/database/context"
	Server "first_api/src/routes"
)

func main() {
	C.ConntectDataBase()
	Server.SetRoutes()
	Server.StartServer(":3300")
}
