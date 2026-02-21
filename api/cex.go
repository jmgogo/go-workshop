package api

import (
	"api_client/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http" // HTTP library
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	currency = strings.ToUpper(currency) // Ensure currency is uppercase for API

	resp, err := http.Get(fmt.Sprintf(apiUrl, currency))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cryptoRate CEXResponse

	if resp.StatusCode == http.StatusOK {
		// Parse response body and return rate
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//  unmarshal the JSON into a struct

		err = json.Unmarshal(bodyBytes, &cryptoRate)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var rate datatypes.Rate
	rate.Currency = currency
	rate.Price = cryptoRate.Bid // Use the bid price from CEXResponse

	return &rate, nil
}
