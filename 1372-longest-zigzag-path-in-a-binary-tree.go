/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    ans := 0

    var helper func(node *TreeNode, right bool, depth int) 
    helper = func(node *TreeNode, right bool, depth int) {
        // fmt.Println(right, depth, ans)
        if depth > ans {
            ans = depth
        }
        if right {
            if node.Left != nil {
                helper(node.Left, false, depth + 1)
            }
            if node.Right != nil {
                helper(node.Right, true, 1)
            }
        } else {
            if node.Right != nil {
                helper(node.Right, true, depth + 1)
            } 
            if node.Left != nil {
                helper(node.Left, false, 1)
            }
        }

    }

    if root.Left != nil {
        helper(root.Left, false, 1)
    }
    if root.Right != nil{
        helper(root.Right, true, 1)
    }
    return ans
    
}
