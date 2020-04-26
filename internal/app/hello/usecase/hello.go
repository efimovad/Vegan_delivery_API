package usecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/hello"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Service struct {
}

func (service *Service) GetHello() *models.Hello {
	h := &models.Hello{Message: "Hello from VDelivery!"}

	return h
}

func NewService() hello.Usecase {
	return &Service{}
}
