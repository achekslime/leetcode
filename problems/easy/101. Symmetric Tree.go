package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	qLeft := []*TreeNode{root.Left}
	qRight := []*TreeNode{root.Right}

	for len(qLeft) != 0 && len(qRight) != 0 {
		currLeft := qLeft[0]
		currRight := qRight[0]
		qLeft = qLeft[1:]
		qRight = qRight[1:]
		if currLeft == nil && currRight == nil {
			continue
		}
		if currLeft == nil || currRight == nil || currLeft.Val != currRight.Val {
			return false
		}
		qLeft = append(qLeft, currLeft.Left, currLeft.Right)
		qRight = append(qRight, currRight.Right, currRight.Left)
	}

	return true
}
