package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(maxLevel *int, maxVal *int, level int, root *TreeNode) {
	if root == nil {
		return
	}
	if level > *maxLevel {
		*maxLevel = level
		*maxVal = root.Val
	}
	find(maxLevel, maxVal, level+1, root.Left)
	find(maxLevel, maxVal, level+1, root.Right)
}

func findBottomLeftValue(root *TreeNode) int {
	maxLevel, maxVal := 0, root.Val
	find(&maxLevel, &maxVal, 0, root)
	return maxVal
}
