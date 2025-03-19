/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
//     Val int
//     Next *ListNode
// }
func removeNthFromEnd(head *ListNode, n int) *ListNode {

    hare := head

    // maintain hare and turtle such that
    // turtle is behind hare by n nodes
    // so when the hare reaches the end of the list
    // 
    // so we can just delete the node after the turtle
    for i:= 0; i < n && hare != nil; i++ {
        if hare == nil {
            return head 
        }
        hare = hare.Next
    }

    var turtle *ListNode = &ListNode{Val: 0, Next: head}
    for hare != nil {
        hare = hare.Next
        turtle = turtle.Next
    }

    if turtle.Next == head {
        return head.Next
    }
    turtle.Next = turtle.Next.Next
    return head
}
