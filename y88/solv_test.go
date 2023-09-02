package y88

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    nums1 := []int{1, 2, 3, 0,0,0}
    nums2 := []int{2,5,6}
    merge(nums1,len(nums1)-len(nums2), nums2, len(nums2))
    for _, num := range nums1 {
        fmt.Println(num)
    }
}
