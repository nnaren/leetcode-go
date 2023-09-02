package  y1654

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    forbidden := []int{1,6,2,14,5,17,4}

    fmt.Println(minimumJumps(forbidden, 16, 9 ,7))

}
