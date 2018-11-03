package solution

type NumArray struct {
	p []int
}

func Constructor(nums []int) NumArray {
	p := make([]int, len(nums)+1)
	for i := 1; i < len(p); i++ {
		p[i] = p[i-1] + nums[i-1]
	}
	return NumArray{
		p: p,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.p[j+1] - this.p[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
