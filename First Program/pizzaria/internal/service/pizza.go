package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidatePrice(pizza *models.Pizza) error {
	if pizza.PRECO < 0 {
		return errors.New("pizza price cannot be less than zero")
	}
	return nil
}
