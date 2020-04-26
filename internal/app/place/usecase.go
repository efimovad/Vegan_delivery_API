package place

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IUsecase interface {
	GetPlaces() ([]models.Place, error)
}