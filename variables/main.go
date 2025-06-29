package main

import (
	"fmt"
	"time"
)

// 1.03 CREATING MULTIPLE VARIABLES USING VAR KEYWORD

// var (
// 	Debug       bool      = false
// 	LogLevel    string    = "info"
// 	startUpTime time.Time = time.Now()
// )

// 1.04 SKIPPING THE TYPE OR VALUE WHEN DECLARING VARIABLES

// var (
// 	Debug       = false
// 	LogLevel    = "info"
// 	startUpTime = time.Now()
// )

func main() {

	// 1.05 IMPLEMENTING A SHORT VARIABLE DECLARATION

	// Debug := false
	// LogLevel := "info"
	// startUpTime := time.Now()

	// DECLARING MULTIPLE VARIABLES WITH A SHORT VARIABLE DECLARATION

	Debug, LogLevel, startUpTime := false, "info", time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)

}
