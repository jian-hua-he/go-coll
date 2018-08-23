package coll

import "testing"

func TestSplice(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	coll = coll.Splice(2)
	if len(coll) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 3)
	}
	arr1 := []interface{}(coll)
	if arr1[0] != 3 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[0], 3)
	}
	if arr1[1] != 4 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[1], 4)
	}
	if arr1[2] != 5 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[2], 5)
	}

	coll = coll.Splice(3)
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 0)
	}
}
