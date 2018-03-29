package main

func twoSum(nums []int, target int) []int {
	var indexNumRecord map[int]int
	for i := 0; i < len(nums); i++ {
		if index, ok := indexNumRecord[target-nums[i]]; ok {
			return []int{i, index}
		}
		indexNumRecord[nums[i]] = i

	}
	return nil

}
