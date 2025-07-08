package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	current := root
	for current != nil {
		if current.Val == val {
			return current
		} else if current.Val < val {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return current
}
