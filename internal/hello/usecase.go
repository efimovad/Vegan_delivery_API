package hello

import "Vegan_delivery_API/internal/models"

type Usecase interface {
	GetHello() *models.Hello
}
