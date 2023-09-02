package y1761

import (
	"testing"
)

func Test_solv(t *testing.T) {

	n := 6
	edges := [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}
	minTrioDegree(n, edges)
}
