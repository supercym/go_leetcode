/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    stack := make([]*TreeNode, 0)
    cur := root
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, node.Val)
        cur = node.Right
        
    }
    return ans
}