package coll

import "testing"

func TestSliceLen(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := coll.SliceLen(4, 6)
	if len(coll) != 10 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 10)
	}
	if len(slice1) != 6 {
		t.Errorf("Data was incorrect, get: %d, want: %d", len(slice1), 6)
	}
	if slice1[0] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[0], 5)
	}
	if slice1[1] != 6 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[1], 6)
	}
	if slice1[2] != 7 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[2], 7)
	}
	if slice1[3] != 8 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[3], 8)
	}
	if slice1[4] != 9 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[4], 9)
	}
	if slice1[5] != 10 {
		t.Errorf("Data was incorrect, get: %d, want: %d", slice1[5], 10)
	}
}
