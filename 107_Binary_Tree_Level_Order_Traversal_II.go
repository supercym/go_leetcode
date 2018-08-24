/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    total_level := [][]int{}
    
    if root == nil {
        return total_level
    }
    
    processed := []*TreeNode{root}
    for len(processed) != 0 {
        tmp := []*TreeNode{}
        level := []int{}
        
        for len(processed) != 0 {
            node := processed[0]
            processed = processed[1:]
            
            level = append(level, node.Val)
            
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        total_level = append(total_level, level)
        processed = tmp
    }
    for i, j := 0, len(total_level)-1; i < j; i, j = i+1, j-1 {
        total_level[i], total_level[j] = total_level[j], total_level[i]
    }
    return total_level
}
