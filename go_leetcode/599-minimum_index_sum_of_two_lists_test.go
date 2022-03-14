package go_leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRestaurant(t *testing.T) {
	list11 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list12 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	result1 := findRestaurant(list11, list12)
	assert.Equal(t, []string{"Shogun"}, result1)

}
