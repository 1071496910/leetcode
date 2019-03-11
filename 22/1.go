package main

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	ret := []string{}
	for i := 0; i < n; i++ {
		ins := generateParenthesis(i)
		outs := generateParenthesis(n - i - 1)
		for _, in := range ins {
			for _, out := range outs {
				tmp := "(" + in + ")" + out
				ret = append(ret, tmp)
			}
		}
	}
	return ret
}
