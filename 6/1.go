/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
package main

func makeNextInervalFunc(row int, numRows int) func() int {

	if row == numRows-1 || row == 0 {
		return func() int {
			return 2 * (numRows - 1)
		}
	}

	state := 0

	return func() int {
		if state == 0 {
			state = 1
			return (numRows - 1 - row) * 2
		} else {
			state = 0
			return row * 2
		}
	}
}

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	res := ""

	for i := 0; i < numRows; i++ {
		nextInerval := makeNextInervalFunc(i, numRows)
		for j := i; j < len(s); j += nextInerval() {
			res += string(s[j])
		}
	}
	return res

}

func main() {
	println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
