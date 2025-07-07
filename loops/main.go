package main

import (
	"fmt"
)

func main() {
	//   0        1         2       3
	// names := []string{"john", "Jessica", "chloe", "Alex"}

	// for i := 4; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// config := map[string]string{
	// 	"debug":    "1",
	// 	"logLevel": "warn",
	// 	"version":  "2.3.1",
	// }

	// for key, value := range config {

	// 	fmt.Println(key, "=", value)
	// }

	// Activity 2.01

	// words := map[string]int{

	// 	"Gonna": 3,
	// 	"You":   3,
	// 	"Give":  2,
	// 	"Never": 1,
	// 	"Up":    4,
	// }

	// mostUsedWord := ""
	// mostUsedWordsCount := 0

	// for word, count := range words {

	// 	if count > mostUsedWordsCount {
	// 		mostUsedWordsCount = count
	// 		mostUsedWord = word

	// 	}

	// }

	// fmt.Println("Most popular word :", mostUsedWord)
	// fmt.Println("with a count of: ", mostUsedWordsCount)

	// Activity 2.02

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}
