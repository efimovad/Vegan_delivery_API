package dishusecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/dish"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Usecase struct {
	repo dish.IRepository
}

func NewDishUsecase(repository dish.IRepository) dish.IUsecase {
	return &Usecase{
		repo: repository,
	}
}

func (u * Usecase) GetDishes(cafeID int64) ([]models.Dish, error) {
	params := models.Params{
		Desc:  false,
		Limit: 10,
		Page:  1,
	}
	return u.repo.List(cafeID, params)
}

func (u * Usecase) GetDish(ID int64) (*models.Dish, error) {
	return u.repo.Find(ID)
}

func (u * Usecase) AddDish(dish models.Dish) error {
	panic("implement me")
}