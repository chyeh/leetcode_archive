package solution

import "sort"

type Interval struct {
	Start int
	End   int
}

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	startTimeSchedule, endTimeSchedule := make([]int, len(intervals)), make([]int, len(intervals))
	for i, interval := range intervals {
		startTimeSchedule[i] = interval.Start
		endTimeSchedule[i] = interval.End
	}
	sort.Ints(startTimeSchedule)
	sort.Ints(endTimeSchedule)

	maxTakenNum := 0
	for si, ei := 0, 0; si < len(intervals) && ei < len(intervals); {
		if startTimeSchedule[si] < endTimeSchedule[ei] {
			si++
		} else {
			ei++
		}
		if si-ei > maxTakenNum {
			maxTakenNum = si - ei
		}
	}
	return maxTakenNum
}
