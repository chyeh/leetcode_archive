package solution

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func overlaps(a, b Interval) bool {
	if (a.Start <= b.Start && a.End >= b.Start) ||
		(a.Start >= b.Start && a.Start <= b.End) ||
		(a.Start >= b.Start && a.End <= b.End) ||
		(a.Start <= b.Start && a.End >= b.End) {
		return true
	}
	return false
}

func merge(intervals []Interval) []Interval {
	for i := 0; i < len(intervals); i++ {
		var overlapHappens bool
		for j := i + 1; j < len(intervals); j++ {
			if overlaps(intervals[i], intervals[j]) {
				overlapHappens = true
				if intervals[i].Start > intervals[j].Start {
					intervals[i].Start = intervals[j].Start
				}
				if intervals[i].End < intervals[j].End {
					intervals[i].End = intervals[j].End
				}
				intervals = append(intervals[0:j], intervals[j+1:]...)
				break
			}
		}
		if overlapHappens {
			i--
		}
	}
	return intervals
}
