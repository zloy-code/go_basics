package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for {

		r := rand.Intn(8)

		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			goto STOP
		}
		fmt.Println(r)
	}

STOP:
	fmt.Println("Goto label reached")
}
