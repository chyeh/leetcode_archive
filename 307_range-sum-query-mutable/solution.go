package solution

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	a := NumArray{
		tree: make([]int, len(nums)+1),
		nums: make([]int, len(nums)),
	}
	for i, n := range nums {
		a.nums[i] = n
		a.add(i+1, n)
	}
	return a
}

func (this *NumArray) Update(i int, val int) {
	d := val - this.nums[i]
	this.nums[i] = val
	this.add(i+1, d)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum(j+1) - this.sum(i)
}

func lsb(i int) int {
	return i & (-i)
}

func (this *NumArray) add(i int, val int) {
	for ; i < len(this.tree); i += lsb(i) {
		this.tree[i] += val
	}
}

func (this *NumArray) sum(i int) int {
	var sum int
	for ; i > 0; i -= lsb(i) {
		sum += this.tree[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
