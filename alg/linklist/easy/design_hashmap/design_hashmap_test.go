package design_hashmap

import "testing"

func TestMyHashMap_Get(t *testing.T) {
	key := 10
	value := 9
	obj := Constructor()
	obj.Put(key, value)
	param_2 := obj.Get(key)
	t.Log(param_2)
	obj.Remove(key)
}
