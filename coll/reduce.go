package coll

// ReduceFunc given two parameters and return a calculated value.
// Those two parameters are accumulator and current value.
// Accumulator accumulates the callback’s return values;
// Current value is the current element in Collection.
type ReduceFunc func(a interface{}, c interface{}) interface{}

// Reduce function reduces the collection to a single value.
func (coll *Collection) Reduce(f ReduceFunc, init interface{}) (r interface{}) {
	l := len(*coll)
	if l == 0 {
		return nil
	}

	if init != nil {
		r = init
	}

	arr := ToInterfaces(*coll)
	for i := 0; i < l; i++ {
		r = f(r, arr[i])
	}

	return r
}
