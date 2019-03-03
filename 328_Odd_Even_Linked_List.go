/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    even := &ListNode{}
    evenHead := even
    node := head
    for {
        even.Next = node.Next
        even = even.Next
        if node.Next != nil {
            node.Next = node.Next.Next
            if node.Next != nil {
                node = node.Next
            } else {
                break
            }
        } else {
            break
        }
    }
    if even != nil {
        even.Next = nil
    }
    node.Next = evenHead.Next
    return head
}