package coll

import "testing"

func TestGet(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	for i, v := range coll {
		if v != coll.Get(i) {
			t.Errorf("Result was incorrect, got: %d, want: %d", v, i+1)
		}
	}
}
