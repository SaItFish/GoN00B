// Package linked_list
// @file: model.go
// @date: 2021/2/1
package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
