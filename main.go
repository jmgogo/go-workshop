package main

import (
	"api_client/api"
	"fmt"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rate for %v: %.2f\n", rate.Currency, rate.Price)
}
