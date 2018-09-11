/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     h := reverseList(head.Next)
//     head.Next.Next = head
//     head.Next = nil
//     return h
// }

// 迭代
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p1 := head
    p2 := head.Next
    p1.Next = nil
    for p2 != nil {
        tmp := p2.Next
        p2.Next = p1
        p1 = p2
        p2 = tmp
    }
    return p1
}
