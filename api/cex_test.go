package api_test

import (
	"api_client/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	// We expect an error because the currency is empty
	if err == nil {
		t.Error("Expected an Error, received nil instead")
	}
}
