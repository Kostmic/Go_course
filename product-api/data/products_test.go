package data

import "testing"

func TestChecksValidation(test *testing.T) {
	product := &Product{
		Name:  "Michaels",
		Price: 1.00,
	}
	err := product.Validate()

	if err != nil {
		test.Fatal(err)
	}
}
