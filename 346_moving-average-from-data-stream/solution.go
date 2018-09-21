package solution

type MovingAverage struct {
	data []int
	sum  int
	size int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data: make([]int, 0),
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.size == 0 {
		return 0.0
	}
	if len(this.data) == this.size {
		this.sum += (val - this.data[0])
		this.data = append(this.data[1:], val)
	} else {
		this.data = append(this.data, val)
		this.sum += val
	}
	return float64(this.sum) / float64(len(this.data))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
