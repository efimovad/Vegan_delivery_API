package dish_repo

import (
	"database/sql"
	"github.com/efimovad/Vegan_delivery_API/internal/app/dish"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type DishRepository struct {
	db *sql.DB
}

func NewDishRepository(db *sql.DB) dish.IRepository {
	return &DishRepository{db}
}

func (r *DishRepository) Create(dish *models.Dish) error {
	return r.db.QueryRow(
		"INSERT INTO dishes (name, cafe, ingredients, calories, weight, cost, image, inStock) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		dish.Name,
		dish.Cafe,
		dish.Ingredients,
		dish.Calories,
		dish.Weight,
		dish.Cost,
		dish.Image,
		dish.InStock,
	).Scan(&dish.ID)
}

func (r *DishRepository) Find(id int64) (*models.Dish, error) {
	newDish := &models.Dish{}
	if err := r.db.QueryRow(
		"SELECT id, name, cafe, ingredients, calories, weight, cost, image, inStock FROM dishes WHERE id = $1",
		id,
	).Scan(
		&newDish.ID,
		&newDish.Name,
		&newDish.Cafe,
		&newDish.Ingredients,
		&newDish.Calories,
		&newDish.Weight,
		&newDish.Cost,
		&newDish.Image,
		&newDish.InStock,
	); err != nil {
		return nil, err
	}
	return newDish, nil
}

func (r *DishRepository) List(cafe int64, params models.Params) ([]models.Dish, error) {
	var dishes []models.Dish
	rows, err := r.db.Query(
		"SELECT id, name, cafe, ingredients, calories, weight, cost, inStock, image "+
			"FROM dishes " +
			"WHERE cafe = $1 " +
			"LIMIT $2 OFFSET CASE WHEN $3 > 0 THEN ($3 - 1) * $2 END;",
			cafe, params.Limit, params.Page,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := models.Dish{}
		err := rows.Scan(&item.ID, &item.Name, &item.Cafe, &item.Ingredients, &item.Calories, &item.Weight,
			&item.Cost, &item.InStock, &item.Image)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, item)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	return dishes, nil
}