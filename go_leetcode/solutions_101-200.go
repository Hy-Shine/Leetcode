package go_leetcode

// No.136
func singleNumber(nums []int) int {
	var x int
	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}
	return x
}

// // No.136-2
func singleNumber2(nums []int) int {
	var m = make(map[int]int, len(nums)/2)
	for i := range nums {
		m[nums[i]]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

// NO.191
func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		if num%2 == 1 {
			cnt++
		}
		num = num >> 1
	}
	return cnt
}

// NO.191-2
func hammingWeight2(num uint32) int {
	cnt := uint32(0)
	for num > 0 {
		cnt += num & 1
		num = num >> 1
	}
	return int(cnt)
}
