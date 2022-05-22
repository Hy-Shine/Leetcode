package go_leetcode

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return addBinary2(a, b)
	}
	return addBinary2(b, a)
}

func addBinary2(str1, str2 string) string {
	byteList := make([]byte, len(str1)+1)
	byteSum := map[int]byte{0: '0', 1: '1', 2: '0', 3: '1'}
	byteMap := map[byte]int{'0': 0, '1': 1}

	ok := false
	for i := len(str1) - 1; i >= 0; i-- {
		index := byteMap[str1[i]]
		if ok {
			index++
		}
		if len(str1)-i <= len(str2) {
			index += byteMap[str2[i+len(str2)-len(str1)]]
		}
		ok = index >= 2
		byteList[len(str1)-i] = byteSum[index]
	}
	if ok {
		byteList[0] = '1'
		return string(byteList)
	}
	return string(byteList[1:])
}
