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

import "fmt"
import "sort"

func main() {
	fmt.Println(threeSum([]int{-3, -2, -1, 0, 0, 1, 1, 2, 2, 3, 3, 3, 4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, -4, -1, -4, -2, -3, 2}))
	fmt.Println(threeSum([]int{-4, -1, -4, 0, 2, -2, -4, -3, 2, -3, 2, 3, 3, -4}))
}

func threeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	return threeSumOnSortedArr(nums)

}

func threeSumOnSortedArr(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	ret := make([][]int, 0)
	for i := 0; i+2 < len(nums); {
		println("i=", i)
		println("-nums[i]", -nums[i])
		midRets := twoSumOnSortedArr(-nums[i], nums[i+1:])
		for _, midRet := range midRets {
			midRet = append(midRet, nums[i])
			sort.Ints(midRet)
			ret = append(ret, midRet)
		}
		for i = i + 1; i < len(nums) && nums[i] == nums[i-1]; i++ {
		}

	}
	return ret
}

func twoSumOnSortedArr(target int, nums []int) [][]int {
	if len(nums) < 2 {
		return nil
	}
	avg := target / 2
	i := findNearestIndex(avg, nums)
	j := i + 1
	ret := make([][]int, 0)
	for i >= 0 && j < len(nums) {
		if nums[i]+nums[j] == target {
			ret = append(ret, []int{nums[i], nums[j]})
			for j = j + 1; j < len(nums) && nums[j] == nums[j-1]; j++ {
				println("j=", j)
			}
		} else if nums[i]+nums[j] > target {
			for i = i - 1; i >= 0 && nums[i] == nums[i+1]; i-- {
			}
		} else {
			for j = j + 1; j < len(nums) && nums[j] == nums[j-1]; j++ {
				println("j=", j)
			}
		}

	}
	return ret
}

func findNearestIndex(target int, nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	mid := (l + r) / 2
	for l < r && l != mid && r != mid {
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			break
		}
		mid = (l + r) / 2
		println("mid:", mid)
	}
	for mid > 0 && nums[mid] == nums[mid-1] {
		mid--
	}
	println("mid:", mid)
	return mid
}
