package coll

type FilterFunc func(v interface{}) bool

type FilterFuncWithIndex func(v interface{}, i int) bool

func (coll Collection) Filter(f FilterFunc) Collection {
	results := make(Collection, 0)
	for _, v := range coll {
		if f(v) {
			results = append(results, v)
		}
	}

	return results
}

func (coll Collection) FilterWithIndex(f FilterFuncWithIndex) Collection {
	results := make(Collection, 0)
	for i, v := range coll {
		if f(v, i) {
			results = append(results, v)
		}
	}

	return results
}
