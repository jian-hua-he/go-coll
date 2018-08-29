package coll

// ReduceFunc ...
type ReduceFunc func(a interface{}, c interface{}) interface{}

// Reduce ...
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
