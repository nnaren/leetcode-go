package y1289

import "math"

func minFallingPathSum(grid [][]int) int {
    n := len(grid)
    firstMinSum, secondMinSum := 0, 0
    firstMinIndex := -1
    for i:= 0; i<n; i++ {
        curFirstMinSum, curSecondMinSum := math.MaxInt,math.MaxInt
        curFirstMinIndex := -1
        for j:=0; j<n; j++ {
            curSum := grid[i][j]
            if j != firstMinIndex {
                curSum += firstMinSum
            } else {
                curSum += secondMinSum
            }

            if curSum < curFirstMinSum {
                curSecondMinSum, curFirstMinSum = curFirstMinSum, curSum
                curFirstMinIndex = j
            } else if curSum < curSecondMinSum {
                curSecondMinSum = curSum
            }
        }
        firstMinSum, secondMinSum = curFirstMinSum, curSecondMinSum
        firstMinIndex = curFirstMinIndex
    }
    return firstMinSum
}

