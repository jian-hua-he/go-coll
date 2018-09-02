package coll

// MapFunc is a function that return new interface{} with given interface{}
type MapFunc func(v interface{}) interface{}

// Map method iterates through the Collection
// and passes each value to the given MapFunc.
// The MapFunc is free to modify the item and return it
func (coll *Collection) Map(f MapFunc) Collection {
	results := make(Collection, 0)
	for _, v := range *coll {
		item := f(v)
		results = append(results, item)
	}

	return results
}
