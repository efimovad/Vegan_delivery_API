package hello

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type Usecase interface {
	GetHello() *models.Hello
}
