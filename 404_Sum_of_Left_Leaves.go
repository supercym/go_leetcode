/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    que := []*TreeNode{root}
    ans := 0
    for len(que) != 0 {
        node := que[0]
        que = que[1:]
        if node.Left != nil {
            if node.Left.Left == nil && node.Left.Right == nil {
                ans += node.Left.Val
            } else {
                que = append(que, node.Left)
            }
        }
        if node.Right != nil {
            que = append(que, node.Right)
        }
    }
    return ans
}