package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a, b := 5, 10

	swap(&a, &b)

	fmt.Println(&a)
	fmt.Println(a == 10, b == 5)

}
