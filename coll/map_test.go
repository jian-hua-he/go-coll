package coll

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5"}

	results := Collect(arr).Map(func(v interface{}) interface{} {
		n, _ := strconv.Atoi(v.(string))
		return n
	})

	for _, v := range results {
		n := reflect.ValueOf(v)
		if n.Kind() != reflect.Int {
			t.Errorf("Results type was incorrect, got: %T, want: %s", v, reflect.Int)
			return
		}
	}
}
