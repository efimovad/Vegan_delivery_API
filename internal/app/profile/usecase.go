package profile

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IUsecase interface {
	GetProfile(id int64) (models.Profile, error)
}
