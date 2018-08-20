package solution

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	if z == 0 {
		return true
	}
	for y != 0 {
		x, y = y, x%y
	}
	return z%x == 0
}
