package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize = 10 * GlobalLimit
const (
	CacheKeyBook = "book_"
	CacheKeyCd   = "cd_"
)

var cache map[string]string

func getCache(key string) string {
	return cache[key]
}

func setCache(key, val string) {

	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val

}

func getBook(isbn string) string {
	return getCache(CacheKeyBook + isbn)
}

func setBook(isbn string, name string) {

	setCache(CacheKeyBook+isbn, name)
}

func getCd(sku string) string {
	return getCache(CacheKeyCd + sku)
}

func setCd(sku string, title string) {
	setCache(CacheKeyCd+sku, title)
}

func main() {

	cache = make(map[string]string)

	setBook("1234-5678", "Get Ready To Go")

	setCd("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book: ", getBook("1234-5678"))

	fmt.Println("CD: ", getCd("1234-5678"))

	fmt.Println(cache)
}
