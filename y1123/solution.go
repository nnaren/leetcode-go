package y1123

import (
	"fmt"
	"leetcode/treenode"
)

var (
	ans    *treenode.TreeNode
	deepth int
)

func lcaDeepestLeaves(root *treenode.TreeNode) *treenode.TreeNode {
	ans = root
	deepth = 0
	findChildrenDeep(root, 0)
	return ans
}

func findChildrenDeep(root *treenode.TreeNode, depth int) int {
	if root == nil {
		return 0
	}
	depth++
	leftDeep := findChildrenDeep(root.Left, depth)
	rightDeep := findChildrenDeep(root.Right, depth)
	if leftDeep == rightDeep && (leftDeep+depth) >= deepth {
		ans = root
		deepth = leftDeep + depth
		fmt.Printf("deep %d, val %d\n", deepth, ans.Val)
	}
	if leftDeep > rightDeep {
		return leftDeep + 1
	} else {
		return rightDeep + 1
	}

}
