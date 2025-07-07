package main

import (
	"fmt"
	"unicode"
)

// exercise 3.01

// Program to measure password complexity  - parol mustajkamligini tekshirish uchun dastur

func passwordChercker(pwd string) bool {

	//Parolni ko'p baytli (UTF-8) belgilar uchun xavfsiz bo'lgan rune turiga aylantiring
	pwR := []rune(pwd)

	// parolni uzunligini tekshiring

	if len(pwR) < 8 {
		return false
	}

	// kriterialarni e'lon qiling

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// ko'p baytli belgilarni tekshiring

	for _, v := range pwR {

		if unicode.IsUpper(v) {

			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {

			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol

}
func main() {
	// noto'g'ri parol uchun test
	if passwordChercker("") {
		fmt.Println("Password is good")
	} else {
		fmt.Println("Password is bad")
	}

	// to'g'ri parol uchun test
	if passwordChercker("This!I5a") {
		fmt.Println("Password is good")
	} else {
		fmt.Println("Password is bad")
	}

}
