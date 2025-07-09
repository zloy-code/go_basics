// Exercise 3.05

/*
Ushbu mashqda biz string e'lon qilamiz va uni ko'p baytli belglar bilan to'diramiz
va ushbu string o'zgaruvchimizdagi barcha belgilarni range tsikli yordamida ekranga chop qilamiz
*/
package main

import (
	"fmt"
)

/*
   ko'p baytli belgilar qatnashgan matnlarni tekshirishning yana bir usuli bu len() funksiyasidan
   foydalanishdir uni quidagi checkStr() funksiyasi orqali tushuntramiz
*/

func checkStr(str string) {

	runes := []rune(str)
	fmt.Println("Bytes: ", len(str))
	fmt.Println("Runes: ", len(runes))
	// faqat 10 ta belgi chop qilish uchun

	fmt.Println(string(str[:10]))
	fmt.Println(string(runes[:10]))
}

func main() {

	logLevel := "デバッグ"

	for index, runeVal := range logLevel {

		fmt.Println(index, string(runeVal))
	}

	username := "Sir_King_Über"

	checkStr(username)

	/*
			   natija:
			           Bytes:  14
		               Runes:  13
		               Sir_King_�
		               Sir_King_Ü

					   Ko'rib turganingizdek belgilar 13 ta lekin baytlar 1 taga farq qilyapti...

	*/

}

/*
natija:
         0 デ
         3 バ
         6 ッ
         9 グ



ko'p baytli belgilar qatnashgan matnni xatolarsiz chop qilish shunday ko'rinishda bo'ladi.


*/
