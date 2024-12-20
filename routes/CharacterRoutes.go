package routes

import (
	"dofus-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, controller *controllers.CharacterController) {
	router.HandleFunc("/GetAllCharacters", controller.GetAllCharacters).Methods(http.MethodGet)
	router.HandleFunc("/AddCharacter", controller.CreateCharacter).Methods(http.MethodPost)
}
