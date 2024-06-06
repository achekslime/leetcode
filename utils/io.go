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

func (reader *ReaderWriter) PrintSlice(slice []int) {
	_, _ = fmt.Fprintln(reader.out, slice)
	defer reader.out.Flush()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (reader *ReaderWriter) ReadLinkedListInt() *ListNode {
	length := reader.ReadInt()
	arr := reader.ReadSliceInt(length)
	head := &ListNode{Val: 0}
	prev := head
	for _, v := range arr {
		curr := &ListNode{Val: v}
		prev.Next = curr
		prev = curr
	}
	return head.Next
}

func (reader *ReaderWriter) PrintLinkedListInt(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d,", head.Val)
		} else {
			fmt.Printf("%d", head.Val)
		}
		head = head.Next
	}
	fmt.Print("]")
	defer reader.out.Flush()
}

func (reader *ReaderWriter) ReadKLinkedListInt() []*ListNode {
	count := reader.ReadInt()
	lists := make([]*ListNode, count)
	for i := 0; i < count; i++ {
		lists[i] = reader.ReadLinkedListInt()
	}
	return lists
}
