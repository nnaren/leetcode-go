package  y523

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    //nums := []int{23, 2, 4, 6, 7}
    //k :=6
    //checkSubarraySum(nums,k)

    ad := 1.0
    for i:=0 ;i<58;i++ {
        ad *= 1.05
    }
    fmt.Println(ad*4)
    at := 59 * 0.05
    fmt.Println(at)
    fmt.Println(5*at)
    fmt.Println(4*at)
    fmt.Println(3*at)
    fmt.Println(2*at)

    at2 := 60 * 0.025
    fmt.Println(5*at2)
    fmt.Println(6*at2)
    fmt.Println(8*at2)
}
