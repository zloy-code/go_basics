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

func countBooks() int {
	count := 0
	for key := range cache {
		if len(key) >= len(CacheKeyBook) && key[:len(CacheKeyBook)] == CacheKeyBook {
			count++
		}
	}
	return count
}

func countCDs() int {
	count := 0
	for key := range cache {
		if len(key) >= len(CacheKeyCd) && key[:len(CacheKeyCd)] == CacheKeyCd {
			count++
		}
	}
	return count
}

func main() {

	cache = make(map[string]string)

	setBook("1234-5678", "Get Ready To Go")

	setCd("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book: ", getBook("1234-5678"))

	fmt.Println("CD: ", getCd("1234-5678"))

	fmt.Println("Books in cache:", countBooks())
	fmt.Println("CDs in cache:", countCDs())
}
