package main

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

/***

布隆过滤器的原理是，当一个元素被加入集合时，通过K个散列函数将这个元素映射成一个位数组中的K个点，
把它们置为1。检索时，
我们只要看看这些点是不是都是1就（大约）知道集合中有没有它了：
如果这些点有任何一个0，则被检元素一定不在；如果都是1，则被检元素很可能在。这就是布隆过滤器的基本思想。
***/

// BloomFilter 布隆过滤器
type BloomFilter struct {
	Bits []byte
	m    uint
	K    uint // 执行哈系函数的次数
	b    *bitset.BitSet
}

// hashData 长度是32
func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1} // to grab another bit of data
	hasher := murmur3.New128()
	hasher.Write(data) // #nosec
	v1, v2 := hasher.Sum128()
	hasher.Write(a1) // #nosec
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}

func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}
func (bf *BloomFilter) location(h [4]uint64, i uint) uint {
	//return uint(location(h, i) % uint64(f.m))
	return  uint(0)
}

// Put 存储数据,操作位
func (bf BloomFilter) Put(val []byte) error {
	h := baseHashes(val)
	for i := uint(0); i < bf.K; i++ {
		bf.b.Set(bf.location(h, i))
	}
	return nil
}

func (bf *BloomFilter) Contain(val []byte) bool {
	h := baseHashes(val)
	for i := uint(0); i < bf.K; i++ {
		if !bf.b.Test(bf.location(h, i)) {
			return false
		}
	}
	return true
}
