package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lrlevel(root *TreeNode) (int, int) {
	cnt := 1
	l, r := root.Left, root.Right
	for r != nil {
		l = l.Left
		r = r.Right
		cnt++
	}
	if l == nil {
		return cnt, cnt
	} else {
		return cnt + 1, cnt
	}

}

func powerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	r := powerOfTwo(n / 2)
	if n%2 == 0 {
		return r * r
	}
	return 2 * r * r
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := lrlevel(root)
	if l == r {
		return powerOfTwo(r) - 1
	}
	ans := 1
	ans += countNodes(root.Left)
	ans += countNodes(root.Right)
	return ans
}
