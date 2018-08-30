package coll

// MapFunc
type MapFunc func(v interface{}) interface{}

// Map will receive a MapFunc and return a new Collection
func (coll *Collection) Map(f MapFunc) Collection {
	results := make(Collection, 0)
	for _, v := range *coll {
		item := f(v)
		results = append(results, item)
	}

	return results
}
