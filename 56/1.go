package main

import (
	"fmt"
	"sort"
)

/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.
*/

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

type IntervalSet []Interval

func (a IntervalSet) Len() int           { return len(a) }
func (a IntervalSet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntervalSet) Less(i, j int) bool { return a[i].Start < a[j].Start }

func merge(intervals []Interval) []Interval {

	intervals = append(make([]Interval, 0, len(intervals)), intervals...) //避免修改原始数据

	if len(intervals) < 2 {
		return intervals
	}

	sort.Sort(IntervalSet(intervals))

	var res []Interval

	tmpInterval := Interval{
		Start: intervals[0].Start,
		End:   intervals[0].End,
	}

	i := 0
	for j := 1; j < len(intervals); j++ {
		if tmpInterval.End >= intervals[j].Start {
			if intervals[j].End >= tmpInterval.End {
				tmpInterval.End = intervals[j].End
			}
		} else {
			res = append(res, tmpInterval)
			i = j
			tmpInterval.Start = intervals[i].Start
			tmpInterval.End = intervals[i].End
		}
	}
	res = append(res, tmpInterval)
	if i == len(intervals) {
		res = append(res, intervals[len(intervals)-1])
	}
	return res

}

func main() {
	b := IntervalSet{
		{
			Start: 1,
			End:   3,
		},
		{
			Start: 8,
			End:   10,
		},
		{
			Start: 2,
			End:   10,
		},
		{
			Start: 15,
			End:   18,
		},
	}
	fmt.Println(b)
	fmt.Println(merge(b))
	fmt.Println(b)

}
