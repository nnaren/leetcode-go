package y207CourseSchedule


func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	for _,info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	visited := make([]int, numCourses)
	var result []int
	valid := true
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v:= range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for i:= 0; i<numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		indeg = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		// 学习某课程的后就对某个课程先修进度加1
		edges[info[1]] = append(edges[info[1]], info[0])
		//
		indeg[info[0]]++
	}

	q := []int{}
	// 入度减为0则可以学习
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			// 入度减一
			indeg[v]--
			// 入度减为0则可以学习
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
