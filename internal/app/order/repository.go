package order

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IRepository interface {
	GetAll(user string, params models.Params) ([]models.Order, error)
	Create(newOrder models.Order) (int64, error)
	UpdateStatus(id int64, status int64) error
}