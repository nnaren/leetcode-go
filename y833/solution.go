package y833


func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    n, m := len(s), len(indices)

    ops := map[int][]int{}
    for i := 0; i < m; i++ {
        ops[indices[i]] = append(ops[i], i)
    }

    ans := ""
    for i := 0; i < n; {
        succeed := false
        _, ok := ops[i]
        if ok {
            for _, pt := range ops[i] {
                if s[i : i + min(len(sources[pt]), n - i)] == sources[pt] {
                    succeed = true
                    ans += targets[pt]
                    i += len(sources[pt])
                    break
                }
            }
        }
        if !succeed {
            ans += string(s[i])
            i++
        }
    }
    return ans
}

func min(x int, y int) int {
    if x > y {
        return y
    }
    return x
}

