/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 先确定链表中点，将链表后半段翻转，比较前后半段是否相同
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    slow := head
    fast := head
    rev := head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast.Next == nil {
        rev = reverseList(slow)
    } else if fast.Next.Next == nil{
        slow = slow.Next
        rev = reverseList(slow)
    }
    p1 := head
    p2 := rev
    for p1 != nil && p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    return true
}

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