/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := make([]int,0)
    recur(root, &res)
    return res
}

func recur(root *TreeNode, res *[]int) {
    if root == nil {
        return  
    }
    recur(root.Left, res) 
    *res = append(*res, root.Val)
    recur(root.Right, res) 
}
