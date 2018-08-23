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

func TestSpliceLen(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	r1 := coll.SpliceLen(2, 2)
	if len(coll) != 5 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 5)
	}
	if len(r1) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r1), 2)
	}
	arr1 := []interface{}(r1)
	if arr1[0] != 3 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr1[0], 3)
	}
	if arr1[1] != 4 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr1[1], 4)
	}

	r2 := coll.SpliceLen(0, 3)
	if len(coll) != 5 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 5)
	}
	if len(r2) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r2), 3)
	}
	arr2 := []interface{}(r2)
	if arr2[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr2[0], 1)
	}
	if arr2[1] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr2[1], 2)
	}
	if arr2[2] != 3 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr2[2], 3)
	}
}
