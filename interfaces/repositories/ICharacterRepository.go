package repositories

import "dofus-api/models"

type ICharacterRepository interface {
	GetAllCharacters() ([]models.Character, error)
	CreateCharacter(models.Character) error
}
