package coll

import "testing"

func TestReduce(t *testing.T) {
	coll1 := Collection{1, 2, 3, 4}

	r1 := coll1.Reduce(func(a interface{}, c interface{}) interface{} {
		return a.(int) + c.(int)
	}, 5)
	if r1 != 15 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1, 15)
	}

	coll2 := Collection{
		Person{
			Name: "R2D2",
			Age:  500,
		},
		Person{
			Name: "3PO",
			Age:  600,
		},
		Person{
			Name: "Darth Vader",
			Age:  50,
		},
	}
	r2 := coll2.Reduce(func(a interface{}, c interface{}) interface{} {
		return a.(int) + c.(Person).Age
	}, 0)
	if r2 != 1150 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r2, 1150)
	}
}
