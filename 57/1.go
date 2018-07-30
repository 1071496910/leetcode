package main

import (
	"fmt"
)

/*

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
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

func lapped(l Interval, r Interval) bool {
	return (l.Start >= r.Start && l.Start <= r.End) || (r.Start >= l.Start && r.Start <= l.End)
}

func merge(l Interval, r Interval) Interval {
	if r.Start < l.Start {
		return sortedMerge(r, l)
	}
	return sortedMerge(l, r)
}

func sortedMerge(l Interval, r Interval) Interval {
	res := Interval{
		Start: l.Start,
		End:   l.End,
	}
	if l.End < r.End {
		res.End = r.End
	}
	return res
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) < 1 {
		return []Interval{
			newInterval,
		}
	}

	intervals = append(make([]Interval, 0, len(intervals)), intervals...) //避免修改原始数据
	res := make([]Interval, 0)

	i := 0
	tmpInterval := Interval{
		Start: newInterval.Start,
		End:   newInterval.End,
	}

	if newInterval.End < intervals[0].Start {
		return append([]Interval{tmpInterval}, intervals...)
	}

	if newInterval.Start > intervals[len(intervals)-1].End {
		return append(intervals, newInterval)
	}

	for ; i < len(intervals); i++ {
		if !lapped(intervals[i], tmpInterval) {

			if intervals[i].Start < tmpInterval.Start {
				res = append(res, intervals[i])
			} else {
				res = append(res, tmpInterval)
				res = append(res, intervals[i:]...)
				return res
			}

		} else {
			tmpInterval = merge(tmpInterval, intervals[i])
		}
	}
	return append(res, tmpInterval)
}

func main() {
	b := []Interval{
		{
			Start: 1,
			End:   3,
		},
		{
			Start: 6,
			End:   9,
		},
	}
	c := insert(b, Interval{
		Start: 2,
		End:   5,
	})
	fmt.Println(c)
}
