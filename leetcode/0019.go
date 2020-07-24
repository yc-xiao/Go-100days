package main

import "fmt"

// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
type ListNode struct {
	Val  int
	Next *ListNode
}

func toSon(node *ListNode, n int) (endIndex int) {
	if n < 1 {
		return 0
	}
	if node.Next != nil {
		endIndex = toSon(node.Next, n) + 1
	} else {
		return 0
	}
	if endIndex == n {
		node.Next = node.Next.Next
	}
	return endIndex
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{0, head}
	toSon(node, n)
	return node.Next
}

func printNode(node *ListNode) {
	if node == nil {
		return
	}
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}

func main() {
	nums := [...]int{1, 2, 3}
	n := 0
	var head = &ListNode{}
	node := head
	for i, num := range nums {
		if i == 0 {
			node.Val = num
		} else {
			node.Next = &ListNode{Val: num}
			node = node.Next
		}
	}
	printNode(head)
	head = removeNthFromEnd(head, n)
	printNode(head)
}
