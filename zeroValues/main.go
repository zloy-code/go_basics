package main

import (
	"fmt"
	"time"
)

func main() {

	var count int

	fmt.Printf("Count : %#v \n", count)
	// output: Count : 0      zero value for not initialized integer by defautl

	var discount float64

	fmt.Printf("Discount:  %#v \n", discount)
	// output: Discount : 0      zero value for not initialized float64 by defautl

	var debug bool

	fmt.Printf("Debug: %#v  \n", debug)
	// output: Debug: false     zero value for not initialized bool by default

	var message string

	fmt.Printf("String: %#v \n ", message)

	// output: String: ""   zero value for not initialized string by default

	var emails []string

	fmt.Printf("Emails: %#v \n ", emails)
	//output: Emails: []string(nil) zero values for not initialized collection of strings by default

	var startTime time.Time

	fmt.Printf("Time: %#v \n", startTime)

}
