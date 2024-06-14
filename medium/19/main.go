package main

import (
	"log"
	"strconv"
	"strings"
	"unsafe"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func countNodes(head *ListNode, n int) int {
	for range n {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		n++
	}
	return n
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	n = countNodes(head, n) - n
	if n == 0 {
		head = head.Next
		return head
	}
	fastPtr := (*ListNode)(unsafe.Pointer(uintptr(unsafe.Pointer(head)) + unsafe.Sizeof(ListNode{})*uintptr(n-1)))
	fastPtr.Next = fastPtr.Next.Next
	return head
}

type test struct {
	head     *ListNode
	n        int
	expected string
}

func formExpected(node *ListNode, n int) string {
	if node == nil || n == 0 {
		return "[]"
	}
	nn := node
	vals := make([]string, 0)
	idx := 0
	for nn != nil {
		vals = append(vals, strconv.Itoa(nn.Val))
		nn = nn.Next
		idx++
	}
	return "[" + strings.Join(vals, ",") + "]"
}

func main() {
	tests := []test{
		//{
		//	head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, // Casual case
		//	n:        2,
		//	expected: "[1,2,3,5]",
		//},
		//{
		//	head:     &ListNode{Val: 1}, // Single element
		//	n:        1,
		//	expected: "[]",
		//},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}}, // Remove the last element
			n:        1,
			expected: "[1]",
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}}, // Remove the first element
			n:        2,
			expected: "[2]",
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, // Remove the middle element
			n:        2,
			expected: "[1,3]",
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}}}}}}}},
			n:        7,
			expected: "[1,2,3,5,6,7,8,9,10]",
		},
	}
	for _, t := range tests {
		res := removeNthFromEnd(t.head, t.n)
		if formExpected(res, t.n) != t.expected {
			log.Printf("Test failed, expected %v, got %v", t.expected, formExpected(res, t.n))
		} else {
			log.Printf("Test passed for %p", t.head)
		}
	}
}
