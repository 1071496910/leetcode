package main

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
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

import "fmt"

import "sync"

//num str map
var (
	romanNumSymbol map[int][]string
	romanSymbols   = []string{"I", "V", "X", "L", "C", "D", "M"}
	once           sync.Once
)

func InitSymbolTable() {
	var base = 1
	var index = 0
	romanNumSymbol = make(map[int][]string)
	for i := 0; i < (len(romanSymbols) / 2); i, base, index = i+1, base*10, index+2 {
		symbolList := make([]string, 10)
		symbolList[0] = " "
		symbolList[1] = romanSymbols[index]
		symbolList[2] = symbolList[1] + romanSymbols[index]
		symbolList[3] = symbolList[2] + romanSymbols[index]
		symbolList[4] = romanSymbols[index] + romanSymbols[index+1]
		symbolList[5] = romanSymbols[index+1]
		symbolList[6] = romanSymbols[index+1] + romanSymbols[index]
		symbolList[7] = symbolList[6] + romanSymbols[index]
		symbolList[8] = symbolList[7] + romanSymbols[index]
		symbolList[9] = romanSymbols[index] + romanSymbols[index+2]
		romanNumSymbol[base] = symbolList
	}
	symbolList := make([]string, 10)
	symbolList[0] = " "
	symbolList[1] = romanSymbols[index]
	symbolList[2] = symbolList[1] + romanSymbols[index]
	symbolList[3] = symbolList[2] + romanSymbols[index]
	romanNumSymbol[base] = symbolList

}

func intToRoman(num int) string {
	once.Do(InitSymbolTable)
	ret := ""
	power10 := 1

	for num != 0 {
		c := num % 10
		num = num / 10
		ret = romanNumSymbol[power10][c] + ret
		power10 = power10 * 10
	}
	return ret

}

func main() {

	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
