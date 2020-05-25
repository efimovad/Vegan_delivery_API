package placeusecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/place"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Usecase struct {
	repo place.IRepository
}

func NewPlaceUsecase(repository place.IRepository) place.IUsecase {
	return &Usecase{
		repo: repository,
	}
}

func (u *Usecase) GetPlaces(params models.Params) ([]models.Place, error) {
	return u.repo.List(params)
}
