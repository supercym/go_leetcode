/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// ***************  Method 1  *******************
    // dummyRoot := &ListNode{Val:0,}
    // var ptr *ListNode
    // ptr = dummyRoot
    // for ; (l1 != nil) && (l2 != nil) ; {
    //     if l1.Val <= l2.Val {
    //         ptr.Next = &ListNode{l1.Val, nil}
    //         ptr = ptr.Next
    //         l1 = l1.Next
    //     } else {
    //         ptr.Next = &ListNode{l2.Val, nil}
    //         ptr = ptr.Next
    //         l2 = l2.Next
    //     }
    // }
    // if l1 == nil {
    //     ptr.Next = l2
    // } else {
    //     ptr.Next = l1
    // }
    // return dummyRoot.Next
	
	// ***************  Method 2  *******************
    tmp := []int{}
    for ; l1 != nil; l1 =l1.Next {
        tmp = append(tmp, l1.Val)
    }
    for ; l2 != nil; l2 =l2.Next {
        tmp = append(tmp, l2.Val)
    }
    sort.Ints(tmp)
    
    dummyRoot := &ListNode{Val:0,}
    var ptr *ListNode = dummyRoot
    for _, v := range tmp {
        ptr.Next = &ListNode{v, nil}
        ptr = ptr.Next
    }
    return dummyRoot.Next
}