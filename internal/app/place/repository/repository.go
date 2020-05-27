package place_repo

import (
	"database/sql"
	"github.com/efimovad/Vegan_delivery_API/internal/app/place"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewPlaceRepository(db *sql.DB) place.IRepository {
	return &PlaceRepository{db}
}

func (r *PlaceRepository) List(params models.Params) ([]models.Place, error) {
	var items []models.Place
	rows, err := r.db.Query(`
			SELECT id, name, minCost, grade, image, latitude, longitude, logo, deliveryTime 
			FROM places 
			LIMIT $1 OFFSET CASE WHEN $2 > 0 THEN ($2 - 1) * $1 END;`,
		 params.Limit, params.Page,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		i := models.Place{}
		err := rows.Scan(&i.ID, &i.Name, &i.MinCost, &i.Grade, &i.Image, &i.Latitude, &i.Longitude, &i.Logo, &i.DeliveryTime)

		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return items, nil
}