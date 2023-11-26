package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1A := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2A := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	resultA := addTwoNumbers(&l1A, &l2A)
	resultA.print()

	l1B := ListNode{Val: 0, Next: nil}
	l2B := ListNode{Val: 0, Next: nil}
	resultB := addTwoNumbers(&l1B, &l2B)
	resultB.print()

	l1C := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}
	l2C := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
	resultC := addTwoNumbers(&l1C, &l2C)
	resultC.print()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := ListNode{Val: 0, Next: nil}
	temp := &ans
	isOverflow := false
	for l1 != nil || l2 != nil {
		sum := 0
		// total
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		if isOverflow {
			sum += 1
		}
		// handle overflow condition
		if sum >= 10 {
			isOverflow = true
			sum = sum % 10
		} else {
			isOverflow = false
		}
		// create node
		temp.Next = &ListNode{Val: sum, Next: nil}
		temp = temp.Next
		// next node
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	// append one node with value one when last one node addition overflow
	if isOverflow {
		temp.Next = &ListNode{Val: 1, Next: nil}
	}

	return ans.Next
}

func (n ListNode) print() {
	for {
		fmt.Print(n.Val)
		if n.Next == nil {
			break
		}
		n = *n.Next
	}

	fmt.Println("\n")
}
