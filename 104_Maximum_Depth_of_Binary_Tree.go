/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left_depth := 1 + maxDepth(root.Left)
    right_depth := 1 + maxDepth(root.Right)
    
    depth := 0
    if left_depth > right_depth {
        depth = left_depth
    } else {
        depth = right_depth
    }
    return depth
}
