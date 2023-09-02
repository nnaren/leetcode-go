package y1267

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    grid := [][]int{
        {1,0},
        {1,1},
    }
    fmt.Println(countServers(grid))
}
