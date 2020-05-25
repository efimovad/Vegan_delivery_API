package dish

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IRepository interface {
	List(cafe int64, params models.Params) ([]models.Dish, error)
	Find(id int64) (*models.Dish, error)
}