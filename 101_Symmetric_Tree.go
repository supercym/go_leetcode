/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
//  *********    recursive    *************
    // if root == nil {
    //     return true
    // }
    // return sym(root.Left, root.Right)
    if root == nil {
        return true
    }
    if root.Left == nil && root.Right != nil {
        return false
    }
    if root.Left != nil && root.Right == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    
    stack1 := []*TreeNode{root.Left}
    stack2 := []*TreeNode{root.Right}
    
    for len(stack1) != 0 && len(stack2) != 0 {
        n1 := stack1[len(stack1)-1]
        n2 := stack2[len(stack2)-1]
        
        stack1 = stack1[:len(stack1)-1]
        stack2 = stack2[:len(stack2)-1]
        
        if n1.Val != n2.Val {
            return false
        }
        
        if n1.Left != nil && n2.Right == nil {
            return false
        }
        if n1.Left == nil && n2.Right != nil {
            return false
        }
        if n1.Right != nil && n2.Left == nil {
            return false
        }
        if n1.Right == nil && n2.Left != nil {
            return false
        }
        if n1.Left != nil && n2.Right != nil {
            stack1 = append(stack1, n1.Left)
            stack2 = append(stack2, n2.Right)
        }
        if n1.Right != nil && n2.Left != nil {
            stack1 = append(stack1, n1.Right)
            stack2 = append(stack2, n2.Left)
        }
        
    }
    return true
    
}

// func sym(Left, Right *TreeNode) bool {
//     if Left == nil && Right == nil {
//         return true
//     }
//     if Left != nil && Right != nil && Left.Val == Right.Val {
//         return sym(Left.Left, Right.Right) && sym(Left.Right, Right.Left)
//     }
//     return false
// }






