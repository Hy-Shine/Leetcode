package go_leetcode

func plusOne(digits []int) []int {
	length := len(digits)
	if digits[length-1] < 9 {
		digits[length-1] = digits[length-1] + 1
		return digits
	}

	bit := true
	for i := length - 1; i >= 0; i-- {
		if bit {
			bit = digits[i]+1 > 9
			if bit {
				digits[i] = digits[i] - 9
			} else {
				digits[i] = digits[i] + 1
			}
		}
	}

	if bit {
		list := make([]int, length+1)
		list[0] = 1
		copy(list[1:], digits)
		return list
	}
	return digits
}

func plusOne2(digits []int) []int {
	length := len(digits)
	if digits[length-1] < 9 {
		digits[length-1] = digits[length-1] + 1
		return digits
	}

	list := make([]int, length+1)
	bit := true
	for i := length - 1; i >= 0; i-- {
		if bit {
			bit = digits[i]+1 > 9
			if bit {
				list[i+1] = digits[i] - 9
			} else {
				list[i+1] = digits[i] + 1
			}
		} else {
			list[i+1] = digits[i]
		}
	}

	if bit {
		list[0] = 1
		return list
	}
	return list[1:]
}
