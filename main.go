package main

import (
	"bufio"
	"leetcode/solution"
	"leetcode/utils"
	"os"
)

func initBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	return in, out
}

func main() {
	console := utils.NewReaderWriter()

	// input
	n := console.ReadInt()

	// output
	ans := solution.ClimbStairs(n)
	console.PrintInt(ans)
}
