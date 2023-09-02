package y1289

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    grid := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println(minFallingPathSum(grid))
}
