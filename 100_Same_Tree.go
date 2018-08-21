/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    tmp1 := p
    tmp2 := q
    if tmp1 == nil && tmp2 != nil {
        return false
    } else if tmp1 != nil && tmp2 == nil {
        return false
    } else if tmp1 == nil && tmp2 == nil {
        return true
    } else if tmp1.Val != tmp2.Val {
        return false
    }
    return isSameTree(tmp1.Left, tmp2.Left) && isSameTree(tmp1.Right, tmp2.Right)
}
