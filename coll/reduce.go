package coll

type ReduceFunc func(c interface{}, i interface{}) interface{}

func (coll *Collection) Reduce(f ReduceFunc, init interface{}) interface{} {
	return nil
}
