package main

import "fmt"

func main() {

	// 	comment1 := `This is the best thing

	// i have ever seen!!!
	// `
	// 	// Bu holatda \n ishlamaydi
	// 	comment2 := `This is the best \n thing i have ever seen`

	// 	// qo'shtirnoq ichida \n ishlaydi
	// 	comment3 := "This is the best \n thing i have ever seen"

	// 	fmt.Print(comment1)
	// 	// farqni ko'rish uchun Println funksiyasidan foydalanildi
	// 	fmt.Println(comment2)
	// 	fmt.Print(comment3)

	// RUNE

	// Ko'rib turganingizdek matnimizda ba'zi noodatiy belgilar bor misol uchun { Ü }
	username := "Sir_King_Über"
	runes := []rune(username)

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}

	fmt.Println(" ")
	// natija:  83 105 114 95 75 105 110 103 95 195 156 98 101 114

	// Bu natija username o'zgaruvchimizning bayt qiymatlari
	// Keling endi uni ortga matn holatiga qaytarishga harakat qilib ko'ramiz

	for i := 0; i < len(username); i++ {

		fmt.Print(string(username[i]), " ")
	}

	fmt.Println(" ")

	// natja:  S i r _ K i n g _ Ã  b e r

	/* Ko'rib turganingizdek natija biz kutgandek emas buning sababi username o'zgaruvchimizdagi
		       { Ü } belgisi 1 baytdan ko'p joy egallashi natijasida compiler uni noto'g'ri tarjima qildi

			   buni to'g'irlash uchun dasturimizga username elon qilishdan keyin quidagi qatorni
			   qo'shamiz :

	            runes := []rune(username)

	   Endi esa dasturdagi loop qismini o'zgartirib ishga tushuramiz

	*/

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}

	/*

		natija : 83 105 114 95 75 105 110 103 95 195 156 98 101 114
	             S i r _ K i n g _ Ã  b e r
	             S i r _ K i n g _ Ü b e r

	    Oradagi farqni ko'rishingiz mumkin

	*/

}

/*
String turiga qo'shimcha sifatida Rune turi mavjud bo'lib
bu bitta UTF-8 ko'p baytli belgini saqlash uchun yetarli xotiraga ega tur hisoblanadi.
Gap shundaki ba'zi  belgilar 1 baytdan ortiq joy egallaydi. Oddiy matnlarda bir belgi 1 bayt
joy egallasa UTF-8 da belgilar 4 baytgacha joy egallaydi shuni hisobga olgan holda Go dasturlash
tili Rune turidan foydalanadi

*/
