package main

import (
	"fmt"
)

/*
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

type numFreq struct {
	Num  int
	Freq int
}

func (nf *numFreq) Less(obj sortabelObj) sortabelObj {
	if nf.Freq < obj.(*numFreq).Freq {
		return nf
	}
	return obj
}

func (nf *numFreq) Equal(obj sortabelObj) bool {
	return true

}

func topKFrequent(nums []int, k int) []int {

	var freMap map[int]int

	for _, v := range nums {
		freMap[v]++
	}

	numFreqs := []*numFreq{}
	for i, v := range freMap {
		numFreqs = append(numFreqs, &numFreq{
			Num:  i,
			Freq: v,
		})
	}

	length := len(numFreqs)
	initHeap(numFreqs)
	for i := 0; i < k; {
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length = length - 1
		shiftDown(nums, 0, length, targetIndexNum)
		fmt.Println(nums)
		if i > 0 && nums[length] == nums[length+1] {
			continue
		}
		i++
	}

	ret := []int{}
	index := len(nums) - 1
	for i := 0; i < k; {
		if i > 0 && ret[i-1] == nums[index] {

		} else {
			ret = append(ret, nums[index])
			i++
		}
		index--

	}
	return ret
}

func initHeap(nums []sortabelObj) {
	length := len(nums)
	if length == 0 || length == 1 {
		return
	}

	length = len(nums)
	for i := (length / 2) - 1; i >= 0; i-- {
		shiftDown(nums, i, length, targetIndexNum)
	}
}

type sortabelObj interface {
	Less(obj sortabelObj) sortabelObj
	Equal(obj sortabelObj) bool
}

func targetIndexNum(nums []sortabelObj, i, j int, length int) (int, sortabelObj) {
	if j >= length {
		return i, nums[i]
	}

	if nums[i].Less(nums[j]).Equal(nums[i]) {
		return i, nums[i]
	}

	return j, nums[j]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type sortFunc func(nums []sortabelObj, i, j int, length int) (int, sortabelObj)

func shiftDown(nums []sortabelObj, i int, length int, sortF sortFunc) {

	needShift := true
	for needShift {
		needShift = false
		left := 2*i + 1
		right := 2*i + 2

		if left >= length && right >= length {
			return
		}
		targetChildIndex, targetChildValue := sortF(nums, left, right, length)

		//if nums[i] < targetChildValue {
		if nums[i].Less(targetChildValue).Equal(targetChildValue) {
			nums[i], nums[targetChildIndex] = nums[targetChildIndex], nums[i]
			needShift = true
			i = targetChildIndex
		}
	}
}

func main() {
	//a := []int{4,234,12,34,5,1,1, 2,  3, 4, 4, 4, 5, 5235,632,87,2,2134,}
	a := []int{4, 1, -1, 2, -1, 2, 3}
	//a := []int{1,1,1,2,2,3}
	fmt.Println(topKFrequent(a, 2))
	return
}
