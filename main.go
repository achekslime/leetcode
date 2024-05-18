package main

import (
	"bufio"
	"fmt"
	"leetcode/solution"
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
	in, out := initBuffers()
	defer out.Flush()

	var n int
	_, _ = fmt.Fscan(in, &n)

	// slice
	// todo wrap in functions
	//arr := make([]int, n)
	//
	//for i := 0; i < n; i++ {
	//	var value int
	//	_, _ = fmt.Fscan(in, &value)
	//	arr[i] = value
	//}

	// string
	// todo wrap in functions
	var s string
	_, _ = fmt.Fscan(in, &s)

	// output
	d := solution.LongestPalindrome(s)
	_, _ = fmt.Fprintln(out, d)
}
