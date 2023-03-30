/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

    ans := []int{}
    var traverse func (root *TreeNode) 
    traverse = func (root *TreeNode) {
        if root == nil {
            return
        }
        traverse(root.Left)
        ans = append(ans, root.Val)
        traverse(root.Right)
        return
    }
    traverse(root)
    return ans
}
