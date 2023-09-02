package y1654

import "fmt"

func minimumJumps(forbidden []int, a int, b int, x int) int {
    lower:=0
    maxVal := 0
    for _, val := range forbidden {
        maxVal = max(maxVal, val)
    }
    upper := max(maxVal+a, x) +b

    q:= [][3] int{[3]int {0,1,0}} // 位置，方向，步数

    visited := make(map[int]bool)
    forbiddenSet := make(map[int]bool)
    visited[0] = true

    for _,position := range forbidden {
        forbiddenSet[position] = true
    }
    for len(q) > 0 {
        // 出队
        position, direction, step := q[0][0], q[0][1], q[0][2]
        q = q[1:]
        if position == x {
            return step
        }
        nextPosition, nextDirection := position +a,1
        _, isVisited := visited[nextPosition * nextDirection]
        _, isForbidden := forbiddenSet[nextPosition]
        if lower <= nextDirection && nextPosition <= upper && !isVisited && !isForbidden {
            fmt.Println("Add: " ,nextPosition ," , ",nextDirection, " ,", step+1)
            visited[nextPosition * nextDirection] = true
            q = append(q, [3]int{nextPosition, nextDirection, step+1})
        }
        if direction == 1 {
            nextPosition, nextDirection := position -b, -1
            _, isVisited := visited[nextPosition * nextDirection]
            _, isForbidden := forbiddenSet[nextPosition]
            if lower <= nextPosition && nextPosition <= upper && !isVisited && !isForbidden {
                fmt.Println("Add: " ,nextPosition ," , ",nextDirection, " ," ,step+1)
                visited[nextPosition * nextDirection] = true
                q = append(q, [3]int{nextPosition, nextDirection, step + 1})
            }
        }
    }

    return -1
}
func max(x int, y int) int {
    if x < y {
        return y
    }
    return x
}
