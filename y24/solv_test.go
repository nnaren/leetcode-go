package y24

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    lists := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
    ans := swapPairs2(lists)
    for i:= ans; i!=nil;i=i.Next {
        fmt.Printf("%d\n",i.Val)
    }
}
