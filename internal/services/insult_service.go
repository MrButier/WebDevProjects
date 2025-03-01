package services

import (
	"InsultAPI/internal/database"
	"InsultAPI/internal/models"
	"errors"
)

func AddInsult(text string) error { return database.AddInsult(text) }

func GetRandomInsult() (models.Insult, error) { return database.GetRandomInsult() }

func GetAllInsults() (models.Insults, error) { return database.GetAllInsults() }

func AddMultipleInsults(insults []string) error {
	for _, insult := range insults {
		err := database.AddInsult(insult)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteInsult(insultID int) error {
	err := database.DeleteInsult(insultID)
	if err != nil {
		if err.Error() == "insult not found" {
			return errors.New("not_found")
		}
		return err
	}
	return nil
}
