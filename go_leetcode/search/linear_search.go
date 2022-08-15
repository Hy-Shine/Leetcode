package search

import "golang.org/x/exp/constraints"

func LinearSearch[T constraints.Ordered](list []T, target T) int {
	for i := range list {
		if list[i] == target {
			return i
		}
	}
	return -1
}
