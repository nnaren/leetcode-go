package y1123

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/treenode"
	"testing"
)

func Test_solv(t *testing.T) {

	root := &treenode.TreeNode{3,
		&treenode.TreeNode{5, &treenode.TreeNode{6, nil, nil}, &treenode.TreeNode{2, &treenode.TreeNode{7, nil, nil}, &treenode.TreeNode{4, nil, nil}}},
		&treenode.TreeNode{1, &treenode.TreeNode{0, nil, nil}, &treenode.TreeNode{8, nil, nil}},
	}
	ser := treenode.Constructor()
	fmt.Println(ser.Serialize(root))
	ans := lcaDeepestLeaves(root)
	fmt.Println(ser.Serialize(ans))
	assert.Equal(t, 3, ans.Val, "相等")
	assert.Equal(t, 4, ans.Val, "相等")
}
