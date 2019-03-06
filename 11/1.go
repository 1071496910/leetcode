package main

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/
/*
思路：尽量保证间距大的情况下，x和y的最小值大。
从两边开始计算maxArea,那边小那边向中间移动，
这样会略过小的一边作为壁的情况，是因为小的一边作为
壁，大的一边向中间移动的情况下，初始情况是最大值，
因为间隔缩小了，但是x和y的最小值只会更小不会更大。
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		x := height[i]
		y := height[j]
		maxArea = max(maxArea, min(x, y)*(j-i))
		if x > y {
			j--
		} else {
			i++
		}
	}
	return maxArea

}
