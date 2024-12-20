package controllers

import (
	// "dofus-api/business/services"
	"dofus-api/interfaces/services"
	"dofus-api/models"
	"encoding/json"
	"net/http"
)

type CharacterController struct {
	characterService services.ICharacterService
}

func NewCharacterController(service services.ICharacterService) *CharacterController {
	return &CharacterController{characterService: service}
}

func (c *CharacterController) GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := c.characterService.GetAllCharacters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(characters)
}

func (c *CharacterController) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.characterService.CreateCharacter(character); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(character)
}
