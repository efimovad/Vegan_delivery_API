package models

type Place struct {
	ID				int64	`json:"id"`
	Name 			string	`json:"name"`
	MinCost 		int		`json:"minCost"`
	Grade			float32	`json:"grade"`
	DeliveryTime	int		`json:"deliveryTime"`
	Image			string	`json:"image"`
	Longitude		float64	`json:"longitude"`
	Latitude		float64	`json:"latitude"`
}
