package main

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

import "sort"

func closest(t int, l int, r int) int {
	lt := l - t
	rt := r - t
	if lt < 0 {
		lt = -lt
	}
	if rt < 0 {
		rt = -rt
	}
	if lt > rt {
		return r
	}
	return l
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	ret := nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {

		if i == 0 || nums[i] != nums[i-1] {

			twoSum, l, r := target-nums[i], i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] == twoSum {
					return target
				} else if nums[l]+nums[r] < twoSum {
					ret = closest(target, ret, nums[r]+nums[i]+nums[l])
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else {
					ret = closest(target, ret, nums[r]+nums[i]+nums[l])
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}

				}
			}
		}
	}

	return ret

}
