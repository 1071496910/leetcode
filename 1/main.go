package main

import "fmt"
import "sort"

func twoSum(nums []int, target int) []int {
	numsRecord := make([]int, len(nums))
	copy(numsRecord, nums)
	sort.Ints(nums)
	fmt.Println("DEBUG: nums:\n", nums)
	midIndex := findHalf(nums, target>>1)
	if midIndex != -1 {
		left := midIndex - 1
		right := midIndex
		for right < len(nums) && left >= 0 {
			res := nums[left] + nums[right]
			if res > target {
				left--
			} else if res < target {
				right++
			} else {
                return findTarget(numsRecord, []int{
                    nums[left],
                    nums[right],
                }
			}
		}

	}

	return nil
}

func findTarget(nums []int, targetPair []int) []int {
    var leftOk bool = false
    var rightOk bool = false
    res := make([]int,2)
	for i := 0; i < len(nums); i++ {
        if leftOk && rightOk {
            return res
        }

		if nums[i] == target[0] && leftOk == false {
            leftOk = true
			res[0] = i
            continue
		}
        if nums[i] == target[1] {
            rightOk = true
            res[1] = i
            continue
        }
	}
	return -1

}

func findHalf(nums []int, half int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > half {
			fmt.Println("DEBUG: half:\n", i)
			return i
		}
	}
	return len(nums) - 1
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
