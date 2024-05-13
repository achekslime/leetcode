package main

import (
	"bufio"
	"fmt"
	"leetcode/problems/easy"
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

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var value int
		_, _ = fmt.Fscan(in, &value)
		arr[i] = value
	}

	_, _ = fmt.Fprintln(out, easy.MajorityElement(arr))
}
