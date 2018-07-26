package main

/*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

func numTrees(n int) int {
	var sum = 0
	if n == 0 || n == 1 {
		return 1
	}
	for i := 1; i <= n; i++ {
		sum += numTrees(i-1) * numTrees(n-i)
	}
	return sum

}
