package go_leetcode

func longestPalindrome(s string) int {
	var asciimap = make(map[rune]int, 0)
	for _, v := range s {
		if _, ok := asciimap[v]; ok {
			asciimap[v] += 1
		} else {
			asciimap[v] = 1
		}
	}

	var cnt, cnt2 int
	for _, v := range asciimap {
		cnt += v
		if v%2 == 1 {
			cnt2++
		}
	}

	if cnt2 == 0 {
		return cnt
	}
	return cnt - cnt2 + 1
}
