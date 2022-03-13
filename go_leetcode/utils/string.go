package utils

func MinStringLength(a, b string) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}
