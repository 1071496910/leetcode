package main

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ret [][]int = make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSum, l, r := -nums[i], i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] == twoSum {
					ret = append(ret, []int{nums[l], nums[r], nums[i]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[l]+nums[r] < twoSum {
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else {
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
