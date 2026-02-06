package main

import (
	"fmt"
	"packages/data"
)

func printData(message string) {
	fmt.Println(message)
	fmt.Println("Today is " + data.Today)
}
