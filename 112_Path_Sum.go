/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, sum int) bool {
    
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil && root.Val == sum {
        return true
    }
    
    hasLeftPath := hasPathSum(root.Left, sum - root.Val)
    hasRightPath := hasPathSum(root.Right, sum - root.Val)
    
    return hasLeftPath || hasRightPath
}
