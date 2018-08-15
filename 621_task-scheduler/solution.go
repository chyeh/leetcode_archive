package solution

func insert(tasksList [][]byte, task byte, taskCnt map[byte]int) [][]byte {
	var tasks []byte
	for i := 0; i < taskCnt[task]; i++ {
		tasks = append(tasks, task)
	}
	tasksList = append(tasksList, tasks)
	for i := len(tasksList) - 1; i > 0; i-- {
		if len(tasksList[i]) > len(tasksList[i-1]) {
			tasksList[i], tasksList[i-1] = tasksList[i-1], tasksList[i]
		}
	}
	return tasksList
}

func leastInterval(tasks []byte, n int) int {
	taskCnt := make(map[byte]int)
	for _, c := range tasks {
		if _, ok := taskCnt[c]; !ok {
			taskCnt[c] = 1
		} else {
			taskCnt[c]++
		}
	}

	var tasksList [][]byte
	for task, _ := range taskCnt {
		tasksList = insert(tasksList, task, taskCnt)
	}

	maxLen := len(tasksList[0])
	i := 0
	for i = 0; i < len(tasksList); i++ {
		if len(tasksList[i]) != maxLen {
			break
		}
	}

	leftPerSlot := n - (i - 1)
	if leftPerSlot <= 0 {
		return len(tasks)
	}

	left := len(tasks) - maxLen*i
	slots := leftPerSlot * (maxLen - 1)
	d := slots - left
	if d <= 0 {
		return len(tasks)
	}
	return ((maxLen - 1) * (n + 1)) + i
}
