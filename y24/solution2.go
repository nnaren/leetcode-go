package y24

// 返回
func swapPairs2(head *ListNode) *ListNode {
    if head ==nil {
        return nil
    }
    second := head.Next
    head.Next =  swapPairs2(second.Next)
    second.Next = head
    return second
}

