package order

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IUsecase interface {
	GetOrders(userID int64) ([]models.Order, error)
	CreateOrder(newOrder models.Order) (int64, error)
	UpdateStatus(id int64, status string) error
}