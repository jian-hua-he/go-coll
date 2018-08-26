package coll

import "fmt"

func (coll *Collection) Contains(f BoolFunc) bool {
	for _, v := range *coll {
		if f(v) {
			fmt.Println(v)
			return true
		}
	}

	return false
}
