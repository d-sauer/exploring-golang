package main

type InventoryStatistics interface {
	MaxCalories() int
	MinCalories() int
	SumTopCarriers(limit int) int
}
