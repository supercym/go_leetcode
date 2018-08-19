/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func deleteDuplicates(head *ListNode) *ListNode {
    tmp := head
    for tmp != nil {
        if tmp.Next == nil {
            break
        } else if tmp.Next.Val == tmp.Val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    return head   
}