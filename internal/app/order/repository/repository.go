package order_repo

import (
	"database/sql"
	"github.com/efimovad/Vegan_delivery_API/internal/app/order"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
	"github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) order.IRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(newOrder models.Order) (int64, error) {
	err :=  r.db.QueryRow(
		`INSERT INTO orders ("user", cafe, date, cost, status, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		newOrder.User,
		newOrder.Cafe,
		newOrder.Date,
		newOrder.Cost,
		newOrder.Status,
		newOrder.Address,
	).Scan(&newOrder.ID)

	txn, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, _ := txn.Prepare(pq.CopyIn("orders_details", "order_id", "dish_id", "count"))

	for _, item := range newOrder.Items {
		_, err := stmt.Exec(newOrder.ID, int64(item.Dish), int64(item.Count))
		if err != nil {
			return 0, err
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		return 0, err
	}
	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	err = txn.Commit()
	if err != nil {
		return 0, err
	}

	return newOrder.ID, err
}

func (r *OrderRepository) GetAll(userID int64, params models.Params) ([]models.Order, error) {
	var orders []models.Order
	var names pq.StringArray
	var counts pq.Int64Array
	var prices pq.Int64Array

	rows, err := r.db.Query(`
			SELECT ord.id, places.name, ord.date, ord.cost, ord.status, ord.address, places.logo, 
			       (SELECT array_agg(dishes.name)
			       	FROM orders_details 
			       	INNER JOIN dishes ON dishes.id = orders_details.dish_id 
			       	WHERE orders_details.order_id = ord.id),
			       (SELECT array_agg(orders_details.count)
			       	FROM orders_details 
			       	INNER JOIN dishes ON dishes.id = orders_details.dish_id 
			       	WHERE orders_details.order_id = ord.id),
			       (SELECT array_agg(dishes.cost)
			       	FROM orders_details 
			       	INNER JOIN dishes ON dishes.id = orders_details.dish_id 
			       	WHERE orders_details.order_id = ord.id) 
			FROM orders ord
			INNER JOIN places ON places.id = ord.cafe
			WHERE ord."user" = $1 
			ORDER BY id 
			LIMIT $2 OFFSET CASE WHEN $3 > 0 THEN ($3 - 1) * $2 END;`,
		userID, params.Limit, params.Page,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		i := models.Order{}
		err := rows.Scan(&i.ID, &i.CafeName, &i.Date, &i.Cost, &i.Status, &i.Address, &i.CafeLogo, &names, &counts, &prices)

		if err != nil {
			return nil, err
		}

		var itemsFull []models.ItemFull
		for j, _ := range names {
			item := models.ItemFull{}
			item.Name = names[j]
			item.Count = counts[j]
			item.Price = int(prices[j])
			itemsFull = append(itemsFull, item)
		}

		i.ItemsFull = itemsFull
		orders = append(orders, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) UpdateStatus(id int64, status int64) error {
	return r.db.QueryRow("UPDATE orders SET status = $1 WHERE id = $2 RETURNING id",
		status, id,
	).Scan(&id)
}
