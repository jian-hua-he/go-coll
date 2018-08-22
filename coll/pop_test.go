package coll

import "testing"

func TestPop(t *testing.T) {
	coll := Collection{1, 2}

	v1 := coll.Pop()
	if v1 != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", v1, 2)
	}
	if len(coll) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, wnat: %d", len(coll), 1)
	}

	v2 := coll.Pop()
	if v2 != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", v2, 1)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, wnat: %d", len(coll), 0)
	}

	v3 := coll.Pop()
	if v3 != nil {
		t.Errorf("Data was incorrect, get: %v, want: nil", v3)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, wnat: %d", len(coll), 1)
	}
}
