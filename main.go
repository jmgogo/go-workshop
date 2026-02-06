package main

import (
	"fmt"
	"functions/utils"
)

func main() {
	foo, bar := utils.AddAndSubtract(3, 5)
	// add 1 to foo and bar
	utils.Increment(&foo)
	utils.Increment(&bar)

	fmt.Println(foo, bar) // 9 -1
}
