package y1334

import (
	"testing"
)

func Test_solv(t *testing.T) {

	edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	findTheCity(5, edges, 2)
}
