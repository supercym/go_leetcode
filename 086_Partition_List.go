/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    lt := &ListNode{}
    ge := &ListNode{}
    ltHead, geHead := lt, ge
    node := head
    for node != nil {
        if node.Val < x {
            lt.Next = node
            lt = lt.Next
        } else {
            ge.Next = node
            ge = ge.Next
        }
        node = node.Next
    }
    ge.Next = nil
    lt.Next = geHead.Next
    return ltHead.Next
}