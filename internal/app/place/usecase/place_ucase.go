package placeusecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/place"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Usecase struct {
}

func NewPlaceUsecase() place.IUsecase {
	return &Usecase{
	}
}

func (u Usecase) GetPlaces() ([]models.Place, error) {
	place1 := models.Place{
		ID:           1,
		Name:         "Nancy Pizza",
		MinCost:      199,
		Grade:        4.9,
		DeliveryTime: 45,
	}
	
	place2 := models.Place{
		ID:           2,
		Name:         "Veganga",
		MinCost:      399,
		Grade:        4.5,
		DeliveryTime: 35,
	}
	
	place3 := models.Place{
		ID:           3,
		Name:         "Raw to go",
		MinCost:      599,
		Grade:        4.3,
		DeliveryTime: 35,
	}

	return []models.Place{place1, place2, place3}, nil
}
