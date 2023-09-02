package y99

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 中序遍历的实现有迭代和递归两种等价的写法，在本方法中提供迭代实现的写法。使用迭代实现中序遍历需要手动维护栈。
func recoverTree(root *TreeNode)  {
    stack := []*TreeNode{}
    var x, y, pred *TreeNode
    for len(stack) >0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pred != nil && root.Val < pred.Val {
            y = root
            if x == nil {
                x = pred
            } else {
                break
            }
        }
        pred = root
        root = root.Right
    }
    x.Val, y.Val = y.Val, x.Val

}

func minOrder(root *TreeNode) {
    stack := []*TreeNode{}
    for root!=nil || len(stack)>0 {
        for root!= nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Println(root.Val)
        root = root.Right
    }
}
