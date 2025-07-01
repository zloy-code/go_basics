package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {

	if input < 0 {
		return errors.New("Input can not be a negative number")
	} else if input > 100 {
		return errors.New("Input can not be over 100")
	} else if input%7 == 0 {
		return errors.New("Input can not be divisible by 7")
	} else {
		return nil
	}
}

func main() {

	number := 14

	// if number%2 == 0 {
	// 	fmt.Println("The Number is Even!")
	// } else {

	// 	fmt.Println("The Number is Odd!")
	// }

	// if number < 0 {
	// 	fmt.Println("Number can not be negative ")

	// } else if number%2 == 0 {
	// 	fmt.Println("The Number is Even")
	// } else {
	// 	fmt.Println("The Number is Odd")
	// }

	if err := validate(number); err != nil {

		fmt.Println(err)
	} else if number%2 == 0 {
		fmt.Println(number, "is Even")
	} else {
		fmt.Println(number, "is Odd")
	}
}
