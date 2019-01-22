/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
Example 3:

Input: "IX"
Output: 9
Example 4:

Input: "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 5:

Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
package main

import "fmt"

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	last := '0'
	ret := 0
	for curIndex := 0; curIndex < len(s); curIndex++ {
		switch s[curIndex] {
		case 'I':
			ret += 1
			last = 'I'
		case 'V':
			switch last {
			case 'I':
				ret += 3
			default:
				ret += 5
			}
			last = 'V'
		case 'X':
			switch last {
			case 'I':
				ret += 8
			default:
				ret += 10
			}
			last = 'X'
		case 'L':
			switch last {
			case 'X':
				ret += 30
			default:
				ret += 50
			}
			last = 'L'
		case 'C':
			switch last {
			case 'X':
				ret += 80
			default:
				ret += 100
			}
			last = 'C'
		case 'D':
			switch last {
			case 'C':
				ret += 300
			default:
				ret += 500
			}
			last = 'D'
		case 'M':
			switch last {
			case 'C':
				ret += 800
			default:
				ret += 100
			}
			last = 'M'
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
