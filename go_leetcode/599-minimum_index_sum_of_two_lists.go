package go_leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	var list1Map = make(map[string]int, len(list1))
	var list2Map = make(map[string]int, len(list2))

	for i, v := range list1 {
		list1Map[v] = i
	}
	for i, v := range list2 {
		list2Map[v] = i
	}

	indexMap := make(map[int][]string, 0)
	min := len(list1) + len(list2)
	for k, v := range list1Map {
		if v2, ok := list2Map[k]; ok {
			indexMap[v+v2] = append(indexMap[v+v2], k)
			if v+v2 < min {
				min = v + v2
			}
		}
	}

	return indexMap[min]
}
