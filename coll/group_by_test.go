package coll

import (
	"fmt"
	"testing"
)

func TestGroupBy(t *testing.T) {
	pets := Collection{
		Pet{
			Owner:    "John Dow",
			Category: "Cat",
			Name:     "Snow Ball",
		},
		Pet{
			Owner:    "Walter White",
			Category: "Dog",
			Name:     "BB-8",
		},
		Pet{
			Owner:    "Jesse Pinkman",
			Category: "Robot",
			Name:     "R2",
		},
		Pet{
			Owner:    "Jesse Pinkman",
			Category: "Bear",
			Name:     "King",
		},
	}

	results := pets.GroupBy("Owner")
	fmt.Println(results)
}
