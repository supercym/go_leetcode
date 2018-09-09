/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    for head.Val == val{
        head = head.Next
        if head == nil {
            return nil
        }
    }
    tmp := head
    for tmp != nil {
        if tmp.Next == nil {
            break
        }
        for tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
            if tmp.Next == nil {
                return head
            }
        }
        tmp = tmp.Next
    }
    return head
}
