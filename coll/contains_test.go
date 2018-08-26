package coll

import "testing"

type Pet struct {
	Owner    string
	Category string
	Name     string
}

func TestContains(t *testing.T) {
	t.Log("Test Contains Start")

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

	r1 := !pets.Contains(func(v interface{}) bool {
		return v.(Pet).Category > "Robot"
	})
	if !r1 {
		t.Errorf("Data was incorrect, got: %v, want: %v", r1, true)
	}

	r2 := !pets.Contains(func(v interface{}) bool {
		return v.(Pet).Name == "BB-8" && v.(Pet).Owner == "Jesse Pinkman"
	})
	if r2 {
		t.Errorf("Data was incorrect, got: %v, wnat: %v", r2, false)
	}

	t.Log("Test Contains Finish")
}
