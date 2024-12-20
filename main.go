package main

import (
	"dofus-api/business/repositories"
	"dofus-api/business/services"
	"dofus-api/config"
	"dofus-api/controllers"
	"dofus-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("error al conectar a la base de datos %v ", err)
	}
	characterRepository := repositories.NewCharacterRepository(db)
	characterService := services.NewCharacterService(characterRepository)
	characterController := controllers.NewCharacterController(characterService)

	router := mux.NewRouter()
	routes.RegisterRoutes(router, characterController)

	log.Println("servidor iniciado en puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
