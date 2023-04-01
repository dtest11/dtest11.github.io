package design_hashset

import "testing"

func TestLinkList_Contains(t *testing.T) {
	//* Your MyHashSet object will be instantiated and called as such:
	key := 1
	obj := Constructor()
	obj.Add(key)
	param3 := obj.Contains(key)
	t.Log(param3)
	obj.Remove(key)
	param3 = obj.Contains(key)
	t.Log(param3)
}
