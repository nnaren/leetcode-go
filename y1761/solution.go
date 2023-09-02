package y1761

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}

	degree := make([]int, n)
	for _, edges := range edges {
		x, y := edges[0]-1, edges[1]-1
		g[x][y], g[y][x] = 1, 1
		degree[x]++
		degree[y]++
	}
	ans := int(1e9)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if g[i][j] != 1 {
				continue
			}
			for k := j + 1; k < n; k++ {
				if g[i][k] == 1 && g[j][k] == 1 {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == 1e9 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
