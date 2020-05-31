package orderusecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/order"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/pkg/errors"
	"time"
)

type Usecase struct {
	repo order.IRepository
}

func NewOrderUsecase(repository order.IRepository) order.IUsecase {
	return &Usecase{
		repo: repository,
	}
}

func (u *Usecase) GetOrders(user string) ([]models.Order, error) {
	params := models.Params{
		Desc:  false,
		Limit: 10,
		Page:  1,
	}
	return u.repo.GetAll(user, params)
}

func (u *Usecase) CreateOrder(newOrder models.Order) (int64, error) {
	newOrder.Status = models.Created
	newOrder.Date = time.Now()
	return u.repo.Create(newOrder)
}

func (u *Usecase) UpdateStatus(id int64, status string) error {
	var statusID int64
	switch status {
	case "canceled":
		statusID = models.Canceled
	case "preparing":
		statusID = models.Preparing
	case "delivering":
		statusID = models.Delivering
	case "delivered":
		statusID = models.Delivered
	default:
		return errors.New("wrong input")
	}
	return u.repo.UpdateStatus(id, statusID)
}
