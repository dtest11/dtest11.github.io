package twopointers

import (
	"sort"
)

/**
https://leetcode.cn/problems/the-latest-time-to-catch-a-bus/solution/by-tsreaper-qcto/

按没有"你"的情况, 时间顺序遍历有意义的时间点, 模拟上车开车事件

对于有空位的车, 记录最晚能坐该车的时间
对于没空位的车, 记录最晚坐该车的乘客

**/

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Slice(buses, func(i, j int) bool {
		return buses[i] < buses[j]
	})

	sort.Slice(passengers, func(i, j int) bool {
		return passengers[i] < passengers[j]
	})
	set := make(map[int]struct{})
	for _, v := range passengers {
		set[v] = struct{}{}
	}
	// 模拟上车,并记录最后一辆车是否还有空位
	cnt := 0 // 记录最后一辆车的容量
	p := 0
	for _, v := range buses {
		cnt = 0
		// 尝试寻找2个人
		for cnt < capacity && p < len(passengers) {
			if passengers[p] <= v {
				p = p + 1 // 这个人可以上车
				cnt = cnt + 1
			} else {
				break
			}
		}
	}
	p = p - 1
	// 查看最后一辆车上的人数是不是满了==cap
	var result = 0
	if cnt != capacity { // 车没有满
		result = buses[len(buses)-1] // 选择最后发车时间
	} else {
		result = passengers[p] // 车满了，选择最后上车的时间
	}
	// 车辆已经满了，那么只能从最后上车的乘客时间往前推
	// 就从最后一位乘客往前找到空的时间点插队
	/**
	//模拟过后从最后往前找
	//如果最后一个公交车容量满了，就从最后一位乘客往前找到空的时间点插队，在这之前的任意一个时间节点都是可以的，因为最后一位乘客都有机会上车，你比他先来的话肯定也能上车
	//如果最后一个公交车没满，在最后一个公交车到达的时间到就可以了
	int ans=c==0?passengers[j]:buses[buses.length-1];
	while(j>=0 && passengers[j--]==ans)   ans--;  //没有空位就往前找，有空位就插队
	*/
	for p >= 0 {
		// 往前找没人到达的时刻
		_, ok := set[result]
		if !ok {
			break
		}
		result--
		p--
	}
	return result

}
