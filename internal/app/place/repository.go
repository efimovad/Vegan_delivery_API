package place

import "github.com/efimovad/Vegan_delivery_API/internal/models"

type IRepository interface {
	List(params models.Params) ([]models.Place, error)
}
