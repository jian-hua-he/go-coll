package coll

import "testing"

func TestPush(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	coll = Push(coll, 6)
	coll = Push(coll, 7)
	coll = Push(coll, 8)

	for i, v := range coll {
		if v != i+1 {
			t.Errorf("Data was incorrect, got: %d, want: %d", v, i+1)
		}
	}
}
