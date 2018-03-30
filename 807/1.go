package main

/*
807. Max Increase to Keep City Skyline
In a 2 dimensional array grid, each value grid[i][j] represents the height of a building located there. We are allowed to increase the height of any number of buildings, by any amount (the amounts can be different for different buildings). Height 0 is considered to be a building as well.

At the end, the "skyline" when viewed from all four directions of the grid, i.e. top, bottom, left, and right, must be the same as the skyline of the original grid. A city's skyline is the outer contour of the rectangles formed by all the buildings when viewed from a distance. See the following example.

What is the maximum total sum that the height of the buildings can be increased?

Example:
Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
Output: 35
Explanation:
The grid is:
[ [3, 0, 8, 4],
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0] ]

The skyline viewed from top or bottom is: [9, 4, 8, 7]
The skyline viewed from left or right is: [8, 7, 9, 3]

The grid after increasing the height of buildings without affecting skylines is:

gridNew = [ [8, 4, 8, 7],
            [7, 4, 7, 7],
            [9, 4, 8, 7],
            [3, 3, 3, 3] ]

Notes:

1 < grid.length = grid[0].length <= 50.
All heights grid[i][j] are in the range [0, 100].
All buildings in grid[i][j] occupy the entire grid cell: that is, they are a 1 x 1 x grid[i][j] rectangular prism.
*/

/*
思路：
1. 先算出两侧视图的轮廓，即同一条线上的最大值
2. 每个点和左右轮廓对应的点的较小值比较，计算差值，然后求和
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) < 2 {
		return 0
	}

	topBottomList, leftRightList := getCastMap(grid)

	sum := 0
	gridLen := len(grid)

	for i := 0; i < gridLen; i++ {
		for j := 0; j < gridLen; j++ {
			sum = sum + (min(topBottomList[j], leftRightList[i]) - grid[i][j])
		}
	}
	return sum

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getCastMap(grid [][]int) ([]int, []int) {

	gridLen := len(grid)
	topBottomList := make([]int, gridLen)
	leftRightList := make([]int, gridLen)

	for i := 0; i < gridLen; i++ {
		for j := 0; j < gridLen; j++ {
			current := grid[i][j]
			if current > topBottomList[j] {
				topBottomList[j] = current
			}
			if current > leftRightList[i] {
				leftRightList[i] = current
			}
		}
	}
	return topBottomList, leftRightList
}
