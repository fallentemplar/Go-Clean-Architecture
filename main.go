package main

import (
	"clean-architecture/config"
	"clean-architecture/infraestructure/datastore"
	"clean-architecture/infraestructure/router"
	"clean-architecture/registry"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
