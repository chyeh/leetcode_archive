package solution

type Interval struct {
	Start int
	End   int
}

func overlaps(a, b Interval) bool {
	if a.End <= b.Start || b.End <= a.Start {
		return false
	}
	return true
}

func canAttendMeetings(intervals []Interval) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if overlaps(intervals[i], intervals[j]) {
				return false
			}
		}
	}
	return true
}
