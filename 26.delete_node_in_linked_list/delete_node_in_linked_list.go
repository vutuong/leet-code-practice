package deleteNode

/*
https://leetcode.com/problems/delete-node-in-a-linked-list/

237. Delete Node in a Linked List
Write a function to delete a node in a singly-linked list.
You will not be given access to the head of the list,
instead you will be given access to the node to be deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
	} else {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}
