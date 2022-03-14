package go_leetcode

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

func hammingWeight2(num uint32) int {
	cnt := uint32(0)
	for num > 0 {
		cnt += num & 1
		num = num >> 1
	}
	return int(cnt)
}
