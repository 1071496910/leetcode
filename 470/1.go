package main

/*
Given a function rand7 which generates a uniform random integer in the range 1 to 7, write a function rand10 which generates a uniform random integer in the range 1 to 10.

Do NOT use system's Math.random().



 Example 1:

 Input: 1
 Output: [7]
 Example 2:

 Input: 2
 Output: [8,4]
 Example 3:

 Input: 3
 Output: [8,1,10]


  Note:

  rand7 is predefined.
  Each testcase has one argument: n, the number of times that rand10 is called.


   Follow up:

   What is the expected value for the number of calls to rand7() function?
   Could you minimize the number of calls to rand7()?

*/
//hint: https://blog.csdn.net/u013630349/article/details/47908633

var cur = 0

func rand10() int {
	cur = cur % 10
	res := (rand7()+cur)%10 + 1
	cur = cur + 1
	return res
}
