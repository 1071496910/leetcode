/*Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.*/
package main

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	initHeap(nums)
	for i := 0; i < k-1; i++ {
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length = length - 1
		shiftDown(nums, 0, length)
		//fmt.Println(nums)
	}
	return nums[0]
}

func initHeap(nums []int) {
	length := len(nums)
	if length == 0 || length == 1 {
		return
	}

	length = len(nums)
	for i := (length / 2) - 1; i >= 0; i-- {
		shiftDown(nums, i, length)
	}
}

func maxIndexNum(nums []int, i, j int, length int) (int, int) {
	if j >= length {
		return i, nums[i]
	}

	if max(nums[i], nums[j]) == nums[i] {
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

func shiftDown(nums []int, i int, length int) {

	needShift := true
	for needShift {
		needShift = false
		left := 2*i + 1
		right := 2*i + 2

		if left >= length && right >= length {
			return
		}
		maxChildIndex, maxChildValue := maxIndexNum(nums, left, right, length)

		if nums[i] < maxChildValue {
			nums[i], nums[maxChildIndex] = nums[maxChildIndex], nums[i]
			needShift = true
			i = maxChildIndex
		}
	}
}
