package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solutin I: Recursion
/*
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		return -1
	}
	if root.Left.Val != root.Val && root.Right.Val != root.Val {
		return min(root.Left.Val, root.Right.Val)
	}
	if root.Left.Val == root.Val && root.Right.Val != root.Val {
		lAns := findSecondMinimumValue(root.Left)
		if lAns == -1 {
			return root.Right.Val
		} else {
			return min(lAns, root.Right.Val)
		}
	}
	if root.Right.Val == root.Val && root.Left.Val != root.Val {
		rAns := findSecondMinimumValue(root.Right)
		if rAns == -1 {
			return root.Left.Val
		} else {
			return min(rAns, root.Left.Val)
		}
	}
	lAns := findSecondMinimumValue(root.Left)
	rAns := findSecondMinimumValue(root.Right)
	if lAns == -1 && rAns == -1 {
		return -1
	}
	if lAns == -1 {
		return rAns
	}
	if rAns == -1 {
		return lAns
	}
	return min(lAns, rAns)
}
*/

// Improved Code
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		return -1
	}

	var l, r int
	if root.Left.Val == root.Val {
		l = findSecondMinimumValue(root.Left)
	} else {
		l = root.Left.Val
	}
	if root.Right.Val == root.Val {
		r = findSecondMinimumValue(root.Right)
	} else {
		r = root.Right.Val
	}
	if l == -1 || r == -1 {
		return max(l, r)
	}
	return min(l, r)
}

// Test Cases:
// [1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1]
