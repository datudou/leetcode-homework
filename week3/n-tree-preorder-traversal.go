/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int { 
    res := make([]int,0)
    traversal(root, &res)
    return res
}

func traversal(root *Node, res *[]int){
    if root == nil {
        return 
    }

    *res = append(*res,root.Val)
    for i:=0; i < len(root.Children) ; i++{
        traversal(root.Children[i],res)
    }
}


