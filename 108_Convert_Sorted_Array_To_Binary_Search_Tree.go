/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func sortedArrayToBST(nums []int) *TreeNode {
    
    if len(nums) == 0 {
        return nil
    } else if len(nums) == 1 {
        return &TreeNode{nums[0], nil, nil}
    }
    
    mid := int((len(nums)-1)/2)
    root := &TreeNode{nums[mid], nil, nil}
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    
    return root
}
