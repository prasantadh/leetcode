/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tortoise := head
    hare := head.Next

    // loop invariant: before each iteration,
    // all nodes up to tortoise is without duplicates
    // and is sorted
    for hare != nil {
        if tortoise.Val == hare.Val {
            hare = hare.Next
        } else {
            tortoise.Next = hare

            tortoise = hare
            hare = tortoise.Next
        }
    }
    tortoise.Next = hare
    return head
}
