package coll

import "testing"

func TestRemove(t *testing.T) {
	coll := Collection{"Foo", "Bar", "Fuzz"}

	coll.Remove(3)
	if len(coll) != 3 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 3)
	}

	coll.Remove(0)
	if len(coll) != 2 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 2)
	}
	if coll.Get(0) != "Bar" {
		t.Errorf("Data was incorredt, get: %s, want: %s", coll.Get(0), "Bar")
	}
	if coll.Get(1) != "Fuzz" {
		t.Errorf("Data was incorredt, get: %s, want: %s", coll.Get(1), "Fuzz")
	}

	coll.Remove(1)
	if len(coll) != 1 {
		t.Errorf("Collection length was incorrect, get: %d, want: %d", len(coll), 1)
	}
	if coll.Get(0) != "Bar" {
		t.Errorf("Data was incorredt, get: %s, want: %s", coll.Get(0), "Bar")
	}
}
