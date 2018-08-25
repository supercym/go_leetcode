/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    if root.Left == nil && root.Right != nil {
        return 1 + minDepth(root.Right)
    }
    
    if root.Left != nil && root.Right == nil {
        return 1 + minDepth(root.Left)
    }
    
    left_depth := minDepth(root.Left)
    right_depth := minDepth(root.Right)
    
    if left_depth < right_depth {
        return 1 + left_depth
    }
    return 1 + right_depth
}
