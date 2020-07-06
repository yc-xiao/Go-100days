package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func num_to_node(num int) *ListNode {
	var node = &ListNode{}
	root_addr := node
	for num > 0 {
		node.Val = num % 10.0
		num /= 10
		if num <= 0 {
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	node.Next = nil
	return root_addr
}

func node_to_num(node *ListNode) int {
	var num int
	for i := 0; node != nil; i++ {
		num_10 := func(i int) int {
			num_10 := 1
			for j := 0; j < i; j++ {
				num_10 *= 10
			}
			return num_10
		}(i)
		num = node.Val*num_10 + num
		node = node.Next
	}
	return num
}

func to_print(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1_num := node_to_num(l1)
	l2_num := node_to_num(l2)
	l3 := num_to_node(l1_num + l2_num)
	return l3
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var val = 0
	l3 := &ListNode{}
	node := l3
	for l1 != nil || l2 != nil {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val >= 10 {
			node.Val = val - 10
			val = 1
		} else {
			node.Val = val
			val = 0
		}
		if l1 == nil && l2 == nil {
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	if val == 1 {
		node.Next = &ListNode{Val: val}
	}
	return l3
}

func main() {
	l1 := num_to_node(5)
	l2 := num_to_node(5)
	l3 := addTwoNumbers(l1, l2)
	to_print(l3)
}
