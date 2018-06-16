package fenwick_tree

type FenwickTree []int

func lsb(i int) int {
	return i & -i
}

func (t FenwickTree) sum(i int) int {
	var sum int
	for ; i > 0; i -= lsb(i) {
		sum += t[i]
	}
	return sum
}

func (t *FenwickTree) add(i int, n int) {
	for ; i < len(*t); i += lsb(i) {
		(*t)[i] += n
	}
}

func New(a []int) *FenwickTree {
	t := FenwickTree(make([]int, len(a)+1))
	for i, n := range a {
		t.add(i+1, n)
	}
	return &t
}
