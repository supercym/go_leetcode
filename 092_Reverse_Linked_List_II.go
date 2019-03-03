/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n || head == nil {
        return head
    }
    
    var pre *ListNode
    cur := head
    var start *ListNode
    var end *ListNode
    c := 1
    for {
        if c <= m {
            if c == m {
                start, end = pre, cur
            }
            pre = cur
            cur = cur.Next
        } else if c <= n {
            tmp := cur.Next
            cur.Next = pre
            pre = cur
            cur = tmp
        } else {
            if start == nil {
                head = pre
            } else {
                start.Next = pre
            }
            end.Next = cur
            break
        }
        c++
    }
    return head
}