package y23

import (
    "fmt"
    "testing"
)

func Test_solv(t *testing.T) {
    lists := []*ListNode {&ListNode{ 1, &ListNode{4, &ListNode{5 ,nil}}},
        &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
        &ListNode{2, &ListNode{6, nil}},
        }
    ans := mergeKLists2(lists)
    for i:= ans; i!=nil;i=i.Next {
        fmt.Printf("%d\n",i.Val)
    }

}
