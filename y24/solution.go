package y24

 type ListNode struct {
    Val int
    Next *ListNode
 }
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    before := dummy
    for before.Next != nil && before.Next.Next != nil {
        slow := before.Next
        fast := slow.Next
        before.Next = fast
        slow.Next = fast.Next
        fast.Next = slow
        before = slow
    }
    return dummy.Next
}

