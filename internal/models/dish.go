package models

type Dish struct {
	ID			int64	`json:"id"`
	Name 		string	`json:"name"`
	Ingredients string	`json:"ingredients"`
	Calories 	int		`json:"calories"`
	Weight 		int		`json:"weight"`
	Cost 		int		`json:"cost"`
	Image 		string	`json:"image"`
}
