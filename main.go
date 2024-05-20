package main

import (
	"leetcode/solution"
	"leetcode/utils"
)

func main() {
	console := utils.NewReaderWriter()

	// input
	n := console.ReadInt()

	// output
	ans := solution.ClimbStairs(n)
	console.PrintInt(ans)
}
