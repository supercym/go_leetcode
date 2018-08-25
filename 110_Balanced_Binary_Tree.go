/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    que := []*TreeNode{root}
    for len(que) != 0 {
        node := que[0]
        que = que[1:]
        
        left_depth := treeDepth(node.Left)
        right_depth := treeDepth(node.Right)
        
        if (left_depth - right_depth) > 1 || (right_depth - left_depth) > 1 {
            return false
        }
        
        if node.Left != nil {
            que = append(que, node.Left)
        }
        if node.Right != nil {
            que = append(que, node.Right)
        }     
    }
    return true
}

func treeDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    left_depth := 1 + treeDepth(root.Left)
    right_depth := 1 + treeDepth(root.Right)
    
    if left_depth > right_depth {
        return left_depth
    }
    return right_depth
}
