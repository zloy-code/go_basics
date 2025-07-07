package main

import (
	"fmt"
	"runtime"
)

// Numbers(raqamlar) ikki qoidaga ko'ra farqlanadi

// 1 Manfiy qiymat saqlay olish yoki olmasligi
// 2 ular saqlay oladigan eng kichik va eng katta qiymatlarga asoslanib

// manfiy qiymat saqlay oladigan raqamlar signed integers deb nomlanadi
// manfiy qiymat saqlay olmaydigan raqamlar unsigned integers deb nomlanadi

// 1. uint8 - manfiy qiymat saqlay olmaydigan barcha 8 - bitli raqamlar to'plami (0 - 255 gacha)
// 2. uint16 - manfiy qiymat saqlay olmaydigan barcha 16 - bitli raqamlar to'plami (0 - 65535 gacha)
// 3. uint32 - manfiy qiymat saqlay olmaydigan barcha 32 - bitli raqamlar to'plami (0 - 4294967295 gacha)
// 4. uint64 - manfiy qiymat saqlay olmaydigan barcha 64 - bitli raqamlar to'plami (0 - 18446744073709551615 gacha)
// 5. int8 - manfiy qiymat saqlay oladigan barcha 8 - bitli raqamlar to'plami (-128 - 127 gacha)
// 6. int16 - manfiy qiymat saqlay oladigan barcha 16 - bitli raqamlar to'plami (-32768 - 32767 gacha)
// 7. int32 - manfiy qiymat saqlay oladigan barcha 32 - bitli raqamlar to'plami (-2147483648	- 2147483647 gacha)
// 8. int64 - manfiy qiymat saqlay oladigan barcha 64 - bitli raqamlar to'plami (-9223372036854775808 - 9223372036854775807 gacha)

// ikki xil maxsus tur ham bor bular:  uint va int
// bularning hajmi nechchi bitli tizim uchun dasturni ishga tushirishingizga bog'liq 32bit yoki 64bit

// shuncha o'zgaruvchi turi nimaga kerak degan savol tug'iladi. Shunchaki int ni o'zidan foydalansa bo'lmaydimi? Albatta bo'ladi faqat ba'zi hollarda xotiradan foydalanishni kamaytirish uchun boshqa turlarga murojaat qilishga to'g'ri keladi. Misol uchun 32bit va 64bit butun sonlar xotiradan turlicha joy egallaydi. Bu holatni visual korish uchun kichik bir dastur yozib ko'ramiz.

func main() {

	// int turi uchun
	// var list []int
	// int8 turi uchun
	var list []int8

	for i := 0; i < 10000000; i++ {

		list = append(list, 100)
	}

	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)

}

// Natija :
// int turi uchun = TotalAlloc (Heap) = 469 MiB
// int8 turi uchun =
