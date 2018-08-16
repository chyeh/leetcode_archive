package solution

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	r := pow(x, n/2)
	if n%2 == 0 {
		return r * r
	}
	return x * r * r
}
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, -n)
}
