package models

import "time"

const (
	Created = iota
	Preparing = iota
	Delivering = iota
	Delivered = iota
	Canceled = iota
)

type Order struct {
	ID			int64		`json:"id"`
	User		int64		`json:"user,omitempty"`
	Cafe		int64		`json:"cafe,omitempty"`
	CafeName	string		`json:"cafeName,omitempty"`
	Items		[]Item 		`json:"items,omitempty" validate:"min=0,max=30"`
	ItemsFull	[]ItemFull	`json:"itemsFull,omitempty"`
	Date		time.Time	`json:"date"`
	Cost		int64		`json:"cost"`
	Status		int64		`json:"status"`
}

type Item struct {
	Dish  int64	`json:"dish"`
	Count int64	`json:"count"`
}

type ItemFull struct {
	Name	string	`json:"name"`
	Count	int64	`json:"count"`
}