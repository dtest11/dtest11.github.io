package arithmetic

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	l, r := 0, x
	ans := -1
	for l < r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid
		}
	}
	return ans
}

func climbStairs(n int) int {
	c := newCache()
	return c.Get(n)
}

type Cache struct {
	data map[int]int
}

func newCache() *Cache {
	return &Cache{data: map[int]int{1: 1, 2: 2}}
}

func (c *Cache) Get(n int) int {
	val, ok := c.data[n]
	if ok {
		return val
	}
	val = c.Get(n-1) + c.Get(n-2)
	c.data[n] = val
	return val
}
