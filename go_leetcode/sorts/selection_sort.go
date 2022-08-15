package sorts

func selectionSort(s []int) []int {
	min := 0
	for i := 0; i < len(s)-1; i++ {
		min = findSmallest(s[i+1:]) + i + 1
		if s[i] > s[min] {
			s[i], s[min] = s[min], s[i]
		}
	}
	return s
}

// get the min value's index of a array
func findSmallest(p []int) int {
	min := p[0]
	n := 0
	for i := 0; i < len(p); i++ {
		if p[i] < min {
			min = p[i]
			n = i
		}
	}
	return n
}
