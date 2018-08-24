package coll

import "testing"

func TestSplice(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	r1 := coll.Splice(2)
	if len(r1) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r1), 3)
	}
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}
	arr1 := []interface{}(r1)
	if arr1[0] != 3 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[0], 3)
	}
	if arr1[1] != 4 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[1], 4)
	}
	if arr1[2] != 5 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr1[2], 5)
	}

	r2 := coll.Splice(0)
	if len(r2) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r2), 2)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 0)
	}
}

func TestSpliceLen(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	// Except r1 = Collection{3, 4}
	r1 := coll.SpliceLen(2, 2)
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

	// Except coll1 = Collection{1, 2, 5}
	if len(coll) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 3)
	}
	coll1 := []interface{}(coll)
	if coll1[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll1[0], 1)
	}
	if coll1[1] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll1[0], 2)
	}
	if coll1[2] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll1[0], 5)
	}

	// Except r2 = Collection{1}
	t.Logf("ORIGIN: %v", coll)
	t.Log("COLL[0:1]")
	r2 := coll.SpliceLen(0, 1)
	t.Logf("R2: %v", r2)
	t.Logf("COLL: %v", coll)
	if len(r2) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r2), 1)
	}
	arr2 := []interface{}(r2)
	if arr2[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", arr2[0], 1)
	}

	// Except coll = Collection{2, 5}
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}
	coll2 := []interface{}(coll)
	if coll2[0] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll2[0], 2)
	}
	if coll2[1] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll2[1], 5)
	}
}
