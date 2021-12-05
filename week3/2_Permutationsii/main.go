package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// O(kn * log(k)) time
// O(log(k)) space
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	m := (l + r) >> 1
	L := mergeSort(lists, l, m)
	R := mergeSort(lists, m+1, r)
	return mergeTwoLists(L, R)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	protect := &ListNode{0, nil}
	prev := protect

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return protect.Next
}
