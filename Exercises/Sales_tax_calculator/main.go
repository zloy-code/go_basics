package main

import (
	"fmt"
)

func main() {

	total := .0
	// Cake
	total += salesTax(0.99, 0.075)
	// Milk
	total += salesTax(2.75, 0.015)
	//Butter
	total += salesTax(0.87, 0.02)

	fmt.Println("Sales tax total: ", total)
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}
