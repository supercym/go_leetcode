/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    ans := []string{}
    if root == nil {
        return ans
    }
    path := ""
    preOrder(root, path, &ans)
    return ans
}

func preOrder(root *TreeNode, path string, ans *[]string) {
    if root.Left == nil && root.Right == nil {
        *ans = append(*ans, path + strconv.Itoa(root.Val))
        return
    }
    if root.Left != nil {
        preOrder(root.Left, path + strconv.Itoa(root.Val) + "->", ans)
    }
    if root.Right != nil {
        preOrder(root.Right, path + strconv.Itoa(root.Val) + "->", ans)
    }
}
