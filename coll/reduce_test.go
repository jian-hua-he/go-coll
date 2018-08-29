package coll

import "testing"

func TestReduce(t *testing.T) {
	coll := Collection{1, 2, 3, 4}

	r1 := coll.Reduce(func(c interface{}, i interface{}) interface{} {
		return c.(int) + i.(int)
	}, 5)

	if r1 != 15 {
		t.Errorf("Data was incorredt, get: %d, want: %d", r1, 15)
	}
}
