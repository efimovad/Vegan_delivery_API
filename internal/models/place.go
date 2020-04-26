package models

type Place struct {
	ID				int64	`json:"id"`
	Name 			string	`json:"name"`
	MinCost 		int		`json:"minCost"`
	Grade			float32	`json:"grade"`
	DeliveryTime	int		`json:"deliveryTime"`
}
