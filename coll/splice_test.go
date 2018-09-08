package coll

import "testing"

func TestSplice(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	r1 := coll.Splice(2)
	if len(r1) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r1), 3)
	}
	if cap(r1) != 3 {
		t.Errorf("Collection capacity was incorrect, get: %d, want: %d", cap(r1), 3)
	}
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}
	if r1[0] != 3 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1[0], 3)
	}
	if r1[1] != 4 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1[1], 4)
	}
	if r1[2] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1[2], 5)
	}

	r2 := coll.Splice(0)
	if len(r2) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r2), 2)
	}
	if cap(r2) != 2 {
		t.Errorf("Collection capacity was incoorect, get: %d, want: %d", cap(r2), 2)
	}
	if len(coll) != 0 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 0)
	}
	if r2[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r2[0], 1)
	}
	if r2[1] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r2[1], 2)
	}
}

func TestSpliceLen(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	// Except r1 = Collection{3, 4}
	r1 := coll.SpliceLen(2, 2)
	if len(r1) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r1), 2)
	}
	if cap(r1) != 2 {
		t.Errorf("Collection capacity was incorrect, get: %d, want: %d", cap(r1), 2)
	}
	if r1[0] != 3 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1[0], 3)
	}
	if r1[1] != 4 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r1[1], 4)
	}

	// Except coll1 = Collection{1, 2, 5}
	if len(coll) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 3)
	}
	if cap(coll) != 3 {
		t.Errorf("Collection capacity was incorrect, get: %d, want: %d", cap(coll), 3)
	}
	if coll[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll[0], 1)
	}
	if coll[1] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll[0], 2)
	}
	if coll[2] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll[0], 5)
	}

	// Except r2 = Collection{1}
	r2 := coll.SpliceLen(0, 1)
	if len(r2) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(r2), 1)
	}
	if cap(r2) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", cap(r2), 1)
	}
	if r2[0] != 1 {
		t.Errorf("Data was incorrect, get: %d, want: %d", r2[0], 1)
	}

	// Except coll = Collection{2, 5}
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}
	if cap(coll) != 2 {
		t.Errorf("Collection capacity was incorrect, get: %d, want: %d", cap(coll), 2)
	}
	if coll[0] != 2 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll[0], 2)
	}
	if coll[1] != 5 {
		t.Errorf("Data was incorrect, get: %d, want: %d", coll[1], 5)
	}
}
