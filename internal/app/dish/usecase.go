package dish

import (
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type IUsecase interface {
	GetDishes(cafeID int64) ([]models.Dish, error)
	GetDish(ID int64) (*models.Dish, error)
	AddDish(dish models.Dish) error
}
