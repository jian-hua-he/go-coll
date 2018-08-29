package coll

import "testing"

func TestShift(t *testing.T) {
	coll := Collection{1, 2, 3}

	v1 := coll.Shift()
	if v1 != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", v1, 1)
	}
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}

	v2 := coll.Shift()
	if v2 != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", v2, 2)
	}
	if len(coll) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 1)
	}

	v3 := coll.Shift()
	if v3 != 3 {
		t.Errorf("Data was incorrect, get: %d, want: %d", v3, 3)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 0)
	}

	v4 := coll.Shift()
	if v4 != nil {
		t.Errorf("Data was incorrect, get: %d, want: nil", v4)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 0)
	}
}
