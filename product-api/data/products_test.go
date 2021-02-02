package data

import "testing"

func TestChecksValidation(test *testing.T) {
	product := &Product{}

	err := product.Validate()

	if err != nil {
		test.Fatal(err)
	}
}
