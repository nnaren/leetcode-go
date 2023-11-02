package y210CourseSchedule2


func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		// 存储有向图
		edges = make([][]int, numCourses)
		// 标记每个节点的状态：0=未搜索，1=搜索中，2=已完成
		visited = make([]int, numCourses)
		// 用数组来模拟栈，下标 0 为栈底，n-1 为栈顶
		result []int
		// 判断有向图中是否有环
		valid bool = true
		dfs func(u int)
	)

	dfs = func(u int) {
		// 将节点标记为「搜索中」
		visited[u] = 1
		// 搜索其相邻节点
		// 只要发现有环，立刻停止搜索
		for _, v := range edges[u] {
			// 如果「未搜索」那么搜索相邻节点
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			// 如果「搜索中」说明找到了环
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		// 将节点标记为「已完成」
		visited[u] = 2
		// 将节点入栈
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	// 每次挑选一个「未搜索」的节点，开始进行深度优先搜索
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}
	// 如果没有环，那么就有拓扑排序
	// 注意下标 0 为栈底，因此需要将数组反序输出
	for i := 0; i < len(result)/2; i ++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}
