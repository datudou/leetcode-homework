/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var res [][]int
func levelOrder(root *Node) [][]int {
    res = make([][]int,0)
    traversal(root, 0 )
    return res
}

func traversal(root *Node, level int) {
    if root == nil {
        return 
    }

    if level >= len(res) {
       res = append(res, []int{}) 
    }

    res[level] = append(res[level], root.Val)
 
    for i:=0; i< len(root.Children);i++{
        traversal(root.Children[i], level+1)
    }
}

