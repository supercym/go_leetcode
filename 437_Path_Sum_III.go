/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }    
    return dsf(root, sum)+pathSum(root.Left, sum)+pathSum(root.Right, sum)
}

func dsf(root *TreeNode, sum int) int {
    res := 0
    if root == nil {
        return res
    } 
    if root.Val == sum {
        res += 1
    }
    res += dsf(root.Left, sum-root.Val)
    res += dsf(root.Right, sum-root.Val)
    return res
}
