package solution

import "math/rand"

type Solution struct {
	cdf []int
}

func Constructor(w []int) Solution {
	var cdf []int
	var sum int
	for _, v := range w {
		sum += v
		cdf = append(cdf, sum-1)
	}
	return Solution{
		cdf: cdf,
	}
}

func (this *Solution) PickIndex() int {
	max := this.cdf[len(this.cdf)-1] + 1
	num := rand.Intn(max)
	l, r := 0, len(this.cdf)-1
	for l <= r {
		m := (l + r) / 2
		if num <= this.cdf[m] {
			if m == 0 || (m > 0 && num > this.cdf[m-1]) {
				return m
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
