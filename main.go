package main

import (
	"api_client/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1) // Increment WaitGroup counter
		go func() {
			/*
				from Go v1.22 onwards,
				for loop variables have per-iteration scope,
				so there is no need to make a copy of the variable for currency.
				In prior version of Go, the currency variable would be shared
				across each go routine thread,
				leading to the final value of currency being evaluated in all Go routines,
				it was a source of common bugs in Go code.
			*/
			getCurrencyRate(currency)
			wg.Done()
		}()
	}
	wg.Wait() // Wait for all goroutines to finish

}

func getCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rate for %v: %.2f\n", rate.Currency, rate.Price)
}
