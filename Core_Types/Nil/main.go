package main

import "fmt"

/*
   Nil - bu ma'lumot turi emas lekin Go dasturlash tilida maxsus qiymat hisoblanadi va map,
   pointer va interfacelarda bo'sh qiymatni bildiradi

*/

func main() {

	var message []string

	if message == nil {
		fmt.Println("error, unexpected nil value")
		return
	}
	fmt.Println(message)

	/*
		    Bu holatda ko'rib turganimizdek ekranda unexpected nil value natijasi chop qilindi.
			Biz message nomli o'zgaruvchi e'lon qildik ammo unga hechqanday qiymat bermadik
			bu holatda uning qiymati nilga teng ya'ni bo'sh.

	*/
}
