package coll

import "testing"

type Person struct {
	Name string
	Age  int
}

func TestFilter(t *testing.T) {
	pList := Collection{
		Person{
			Name: "Foo",
			Age:  11,
		},
		Person{
			Name: "Bar",
			Age:  2,
		},
		Person{
			Name: "Fuzz",
			Age:  27,
		},
	}
	results := pList.Filter(func(p interface{}) bool {
		return p.(Person).Age > 10
	})

	if len(results) != 2 {
		t.Errorf("Results length was incorrect, got: %d, want: %d", len(results), 2)
	}

	if results[0].(Person).Name != "Foo" {
		t.Errorf("Data was incorrect, got: %s, want: %s", results[0].(Person).Name, "Foo")
	}

	if results[0].(Person).Age != 11 {
		t.Errorf("Results length was incorrect, got: %d, want: %d", results[0].(Person).Age, 11)
	}

	if results[1].(Person).Name != "Fuzz" {
		t.Errorf("Results length was incorrect, got: %s, want: %s", results[1].(Person).Name, "Fuzz")
	}

	if results[1].(Person).Age != 27 {
		t.Errorf("Results length was incorrect, got: %d, want: %d", results[1].(Person).Age, 27)
	}
}
