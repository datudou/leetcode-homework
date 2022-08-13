/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */ 
var valToIndex map[int]int
func buildTree(inorder []int, postorder []int) *TreeNode { 
    valToIndex = make(map[int]int,0)
    for i,v := range inorder {
        valToIndex[v] = i
    }
    return build(inorder, 0,len(inorder)-1, postorder,0 ,len(postorder)-1)
}



func build(inorder []int, inStart, inEnd int, postorder []int, postStart,postEnd int) *TreeNode{
    if inStart  > inEnd{
        return nil
    }
    rootVal := postorder[postEnd]
    index := valToIndex[rootVal]
    leftSize := index - inStart
    root := &TreeNode{
        Val: rootVal,
    }
    root.Left = build(inorder,inStart, index-1,postorder,postStart,postStart+leftSize-1)
    root.Right = build(inorder,index+1,inEnd,postorder, postStart+leftSize, postEnd - 1)
    return root

}
