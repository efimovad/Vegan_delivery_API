package profile_usecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/profile"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Usecase struct {

}

func NewProfileUsecase() profile.IUsecase {
	return &Usecase{}
}

func (u *Usecase) GetProfile(id int64) (models.Profile, error) {
	return models.Profile{
		ID:      id,
		Name:    "Nozim",
		Surname: "Yunusov",
		Phone:   "88888888888",
		Address: "МОСКВА",
	}, nil
}
