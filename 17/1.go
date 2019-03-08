package main

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

var digitLettersMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return letterCombinations2(digits[:1], digits[1:])
}

func letterCombinations2(digit string, digits string) []string {

	ret := []string{}
	midRet := []string{""}

	if len(digits) == 0 {
		if tmp, ok := digitLettersMap[digit]; ok {
			midRet = tmp
		}
		for _, s := range midRet {
			ret = append(ret, s)
		}
		return ret
	}

	tails := letterCombinations2(digits[:1], digits[1:])

	if tmp, ok := digitLettersMap[digit]; ok {
		midRet = tmp
	}

	for _, tail := range tails {
		for _, s := range midRet {
			ret = append(ret, s+tail)
		}
	}
	return ret
}
