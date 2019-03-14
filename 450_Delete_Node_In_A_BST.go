/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == key {
        left, right := root.Left, root.Right
        if left == nil {
            root = right
        } else if right == nil {
            root = left
        } else {
            miniNode := right
            for right.Left != nil {
                right = right.Left
            }
            miniNode = right
            miniNode.Right = deleteNode(root.Right, miniNode.Val)
            miniNode.Left = left
            root = miniNode
        }
    } else if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else {
        root.Left = deleteNode(root.Left, key)
    }
    return root
}