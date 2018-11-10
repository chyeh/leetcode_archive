package solution

// Solution I
/*
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var neg bool
	if x < 0 {
		neg = true
		x = -x
	}
	var number int
	for n := x; n != 0; n = n / 10 {
		d := n % 10
		number = (number * 10) + d
	}
	if number > 2147483647 {
		return 0
	}
	if neg {
		return -number
	}
	return number
}
*/

// Solution II
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var number int
	for n := x; n != 0; n = n / 10 {
		d := n % 10
		number = (number * 10) + d
	}
	if number > 2147483647 || number < -2147483648 {
		return 0
	}
	return number
}
