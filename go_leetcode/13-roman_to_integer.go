package go_leetcode

import "strings"

// Not finish!!!

func romanToInt(s string) int {
	n1 := strings.Count(s, "IV")
	s = strings.ReplaceAll(s, "IV", "")
	n2 := strings.Count(s, "XI")
	s = strings.ReplaceAll(s, "XI", "")

	return n1 + n2
}

var romanMap = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"XI": 9,
	"XL": 40,
	"CX": 90,
	"CD": 400,
	"MC": 900,
}
