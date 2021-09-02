package strings

import (
	"strings"
)

func addBinary(a string, b string) string {
	lena, lenb := len(a), len(b)
	if lena >= lenb {
		return addBinary2(a, b, lena, lenb)
	}
	return addBinary2(b, a, lenb, lena)
}

func addBinary2(s1, s2 string, length1, length2 int) string {
	stringMap := map[int]string{0: "0", 1: "1", 2: "0", 3: "1"}
	j := length2 - 1
	newString, split, next := "", "", ""
	for i := length1 - 1; i >= 0; i-- {
		if j >= 0 {
			split = string(s1[i]) + string(s2[j]) + next
		} else {
			split = string(s1[i]) + next
		}
		index := strings.Count(split, "1")
		newString = stringMap[index] + newString
		if index >= 2 {
			next = "1"
		} else {
			next = ""
		}
		if i == 0 && next == "1" {
			newString = "1" + newString
		}
		j--
	}
	return newString
}
