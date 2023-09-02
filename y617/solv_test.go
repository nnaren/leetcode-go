package y617

import (
    "testing"
)

func Test_solv(t *testing.T) {
    root1 := &TreeNode{1, &TreeNode{3, &TreeNode{5,nil, nil},nil}, &TreeNode{2,nil,nil}}
    root2 := &TreeNode{2, nil, nil}
    to := mergeTrees(root1, root2)
    println(to)

}
