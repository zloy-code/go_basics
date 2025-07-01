package main

import (
	"fmt"
	"time"
)

func main() {

	birthDay := time.Tuesday

	switch birthDay {
	case time.Monday:
		fmt.Println("Monday's child has a beatifull face")
	case time.Tuesday:
		fmt.Println("Tuesday's child is full of grace ")
	case time.Wednesday:
		fmt.Println("Wednesday's child is full of woe")
	case time.Thursday:
		fmt.Println("Thursday's child has far to go")
	case time.Friday:
		fmt.Println("Friday's child is loving and giving")
	case time.Saturday:
		fmt.Println("Saturday's child works hard for living")
	case time.Sunday:
		fmt.Println("Sunday's child is bonny and blithe")
	default:
		fmt.Println("Error, day is not valid")
	}
}
