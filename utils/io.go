package utils

import (
	"bufio"
	"fmt"
	"os"
)

type ReaderWriter struct {
	in  *bufio.Reader
	out *bufio.Writer
}

func NewReaderWriter() *ReaderWriter {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	out.Flush()
	reader := &ReaderWriter{in: in, out: out}
	return reader
}

func (reader *ReaderWriter) ReadInt() int {
	var n int
	_, _ = fmt.Fscan(reader.in, &n)
	return n
}

func (reader *ReaderWriter) ReadString() string {
	var s string
	_, _ = fmt.Fscan(reader.in, &s)
	return s
}

func (reader *ReaderWriter) ReadSliceInt(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		var value int
		_, _ = fmt.Fscan(reader.in, &value)
		arr[i] = value
	}
	return arr
}

func (reader *ReaderWriter) PrintInt(n int) {
	_, _ = fmt.Fprintln(reader.out, n)
	defer reader.out.Flush()
}

func (reader *ReaderWriter) PrintBool(is bool) {
	_, _ = fmt.Fprintln(reader.out, is)
	defer reader.out.Flush()
}
