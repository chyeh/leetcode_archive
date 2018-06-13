package solution

func swap(a [][]int, x int, y int) [][]int {
	var tmp []int
	tmp = a[x]
	a[x] = a[y]
	a[y] = tmp
	return a
}

func insertPerson(a [][]int, index int) [][]int {
	person := a[index]
	a = append(a[0:index], a[index+1:]...)
	left := a[0:person[1]]
	right := a[person[1]:]
	ans := append([][]int{person}, right...)
	ans = append(left, ans...)
	return ans
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) < 2 {
		return people
	}
	for i := 0; i < len(people); i++ {
		for j := i + 1; j < len(people); j++ {
			if people[j][0] == people[i][0] && people[j][1] < people[i][1] {
				swap(people, j, i)
			}
			if people[j][0] > people[i][0] {
				swap(people, j, i)
			}
		}
	}
	for i := 1; i < len(people); i++ {
		people = insertPerson(people, i)
	}
	return people
}
