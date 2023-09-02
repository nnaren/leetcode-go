package y99

import (
    "testing"
)

func Test_solv(t *testing.T) {
    root := &TreeNode{3, &TreeNode{1,nil,nil}, &TreeNode{4, &TreeNode{2, nil,nil}, nil}}
    minOrder(root)
    recoverTree(root)

}
