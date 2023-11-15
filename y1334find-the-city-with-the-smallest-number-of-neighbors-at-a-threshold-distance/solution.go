package y1334

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	minCity := []int{math.MaxInt32, -1}
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = math.MaxInt32
		}
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		dis[from][to], dis[to][from] = weight, weight
	}
	for k := 0; k < n; k++ {
		dis[k][k] = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= minCity[0] {
			minCity[0], minCity[1] = cnt, i
		}
	}

	return minCity[1]
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
