package y23

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
    Val int
    Next *ListNode
}

// 分治
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    dummy :=  &ListNode{} // 哨兵节点
    cur := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            cur.Next = list1
            list1 = list1.Next
        } else {
            cur.Next = list2
            list2 = list2.Next
        }
        cur = cur.Next
    }
    if list1 != nil {
        cur.Next = list1
    }
    if list2 != nil {
        cur.Next = list2
    }
    return  dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    m := len(lists)
    if m == 0 {
        return nil
    }
    if m == 1 {
        return lists[0]
    }
    left := mergeKLists(lists[:m/2])
    right := mergeKLists(lists[m/2:])
    return mergeTwoLists(left, right)
}

