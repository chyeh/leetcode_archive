package solution

// Solution I: Time complexity O(N^2).
/*
func candy(ratings []int) int {
	candyNumArr := make([]int, len(ratings))
	for i := range ratings {
		candyNumArr[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			candyNumArr[i] = candyNumArr[i-1] + 1
		} else {
			for j := i; j > 0; j-- {
				if ratings[j-1] > ratings[j] && candyNumArr[j-1] <= candyNumArr[j] {
					candyNumArr[j-1] = candyNumArr[j] + 1
				}
			}
		}
	}
	var sum int
	for _, candyNum := range candyNumArr {
		sum += candyNum
	}
	return sum
}
*/

// Solution II: Time complexity O(N).
func candy(ratings []int) int {
	candyNumArr := make([]int, len(ratings))
	for i := range ratings {
		candyNumArr[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			candyNumArr[i] = candyNumArr[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candyNumArr[i-1] <= candyNumArr[i] {
			candyNumArr[i-1] = candyNumArr[i] + 1
		}
	}
	var sum int
	for _, candyNum := range candyNumArr {
		sum += candyNum
	}
	return sum
}
