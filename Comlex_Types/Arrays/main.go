package main

import "fmt"

// func defineArray() [10]int {

// 	var arr [10]int

// 	return arr
// }

// func compareArr() (bool, bool, bool) {

// 	var arr1 [10]int

// 	arr2 := [10]int{0}

// 	arr3 := [10]int{0, 0, 0, 0, 0}
// 	arr4 := [10]int{0, 0, 0, 0, 9}

// 	return arr1 == arr2, arr1 == arr3, arr1 == arr4
// }

func sayMyName() string {
	strings := [...]string{
		"name",
		"is",
		"Abbosbek",
		"my",
	}
	// Sprintln funksiyasi yagona string qilib return qilish uchun ishlatildi
	// buyerda elementlar joylashgan indexlar so'zdagi ma'noni chalkashtiradi biz esa uni indexlarni to'g'ri joylashtirish orqali tushunarli holatga olib keldik

	strings[3] = "this"
	return fmt.Sprintln(strings[3], strings[0], strings[1], strings[2])
}

func message() string {

	m := ""

	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {

		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m

}

func fillArray(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr

}

func opArray(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {

		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {

	// comp1, comp2, comp3 := compareArr()

	// Index orqali qiymat berish
	arry := [5]int{0: 11, 3: 9, 2: 10}

	fmt.Println(arry)

	// fmt.Println("[10]int == [10]int{0}", comp1)
	// fmt.Println("[10]int == [10]int{0,0,0,0,0}", comp2)
	// fmt.Println("[10]int == [10]int{0,0,0,0,9}", comp3)

	// fmt.Printf("%#v", defineArray())

	fmt.Println(sayMyName())
	fmt.Println(message())

	var numbers [10]int

	numbers = fillArray(numbers)
	numbers = opArray(numbers)
	fmt.Println(numbers)

}
