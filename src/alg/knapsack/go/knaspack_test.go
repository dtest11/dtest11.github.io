package kanspack

import "testing"

func TestK01(t *testing.T) {
	var data []Payload
	for i := 0; i < 10; i++ {
		data = append(data, Payload{
			Value:  i,
			Weight: i,
		})
	}
	t.Log(K01(data, 10))
}
