package coll

type SortFunc func(a interface{}, b interface{}) bool

func (coll *Collection) Sort(f SortFunc) {
	return
}
