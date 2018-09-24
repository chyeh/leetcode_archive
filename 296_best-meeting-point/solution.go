package solution

func totalDistance(positions []int) int {
	var sum int
	for l, r := 0, len(positions)-1; l < r; l, r = l+1, r-1 {
		sum += positions[r] - positions[l]
	}
	return sum
}

func minTotalDistance(grid [][]int) int {
	var rowPositions []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				rowPositions = append(rowPositions, i)
			}
		}
	}
	var colPositions []int
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 1 {
				colPositions = append(colPositions, j)
			}
		}
	}
	return totalDistance(rowPositions) + totalDistance(colPositions)
}
