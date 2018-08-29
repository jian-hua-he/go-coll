package coll

import "testing"

func TestContains(t *testing.T) {
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
	}

	r1 := pets.Contains(func(v interface{}) bool {
		return v.(Pet).Category == "Robot"
	})
	if !r1 {
		t.Errorf("Data was incorrect, got: %v, want: %v", r1, true)
	}

	r2 := pets.Contains(func(v interface{}) bool {
		return v.(Pet).Name == "BB-8" && v.(Pet).Owner == "Jesse Pinkman"
	})
	if r2 {
		t.Errorf("Data was incorrect, got: %v, wnat: %v", r2, false)
	}
}
