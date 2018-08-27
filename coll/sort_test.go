package coll

import "testing"

func TestSort(t *testing.T) {
	coll := Collection{1, 2, 3, 4, 5}

	coll.Sort(func(a interface{}, b interface{}) bool {
		return a.(int) > b.(int)
	})

	arr := ToInterfaces(coll)
	if arr[0] != 5 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr[0], 5)
	}
	if arr[1] != 4 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr[1], 4)
	}
	if arr[2] != 3 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr[2], 3)
	}
	if arr[3] != 2 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr[3], 2)
	}
	if arr[4] != 1 {
		t.Errorf("Data was incorredt, get: %d, want: %d", arr[4], 1)
	}
}
