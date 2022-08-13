var inOrderValToIndex map[int]int
func buildTree(preorder []int, inorder []int) *TreeNode {
    if preorder == nil || inorder == nil {
        return nil 
    } 
    inOrderValToIndex = make(map[int]int,len(inorder))
    for i, v := range inorder {
        inOrderValToIndex[v] = i
    }

    return build(preorder, 0,len(preorder)-1,inorder,0 ,len(inorder) -1)
}

func build(preorder []int, preStart, preEnd int,inorder []int, inStart,inEnd int) *TreeNode{
    if preStart > preEnd {
        return nil
    }

    root := &TreeNode{
        Val: preorder[preStart],
    }
    
    inorderRootIndex := inOrderValToIndex[root.Val]
    leftLen := inorderRootIndex - inStart
    
    left := build(preorder,preStart+1,preStart+leftLen, inorder,inStart, inorderRootIndex -1)
    right := build(preorder,preStart+leftLen+1, preEnd, inorder,inorderRootIndex+1,inEnd)
    root.Left = left
    root.Right = right
    return root
}

