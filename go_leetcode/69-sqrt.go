package go_leetcode

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	var mid int
	low, high := 0, x>>1
	for low <= high {
		mid = (low + high) >> 1
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			break
		} else if mid*mid < x {
			low++
		} else {
			high--
		}
	}
	return mid
}
