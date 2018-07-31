/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

func lengthOfLongestSubstring(s string) int {
	println(s)

	if len(s) < 2 {
		return len(s)
	}

	maxLength := 1
	charScopeFirstIndexRecord := make(map[byte]int)
	charScopeFirstIndexRecord[s[0]] = 0
	l, r := 0, 1
	for ; r < len(s); r++ {
		println("l=", l, "r=", r)
		if _, ok := charScopeFirstIndexRecord[s[r]]; ok {
			curLen := r - l
			if curLen > maxLength {
				maxLength = curLen
			}
			nli := charScopeFirstIndexRecord[s[r]] + 1
			for l < nli {
				delete(charScopeFirstIndexRecord, s[l])
				l++
			}
		}
		charScopeFirstIndexRecord[s[r]] = r

	}
	if r-l > maxLength {
		return r - l
	}
	return maxLength

}

func main() {
	print(lengthOfLongestSubstring("pwwkew"))

}
