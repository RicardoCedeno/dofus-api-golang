package services

import "dofus-api/models"

type ICharacterService interface {
	GetAllCharacters() ([]models.Character, error)
	CreateCharacter(models.Character) error
}
