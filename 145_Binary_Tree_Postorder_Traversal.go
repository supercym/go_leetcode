/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, node.Val)
        
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    
    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    
    return ans
}