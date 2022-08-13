func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil || list2 ==  nil {
        return nil
    }

    p1 := list1
    p2 := list2
    dumy := &ListNode{}
    p := dumy
    for  p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            p.Next = p1
            p1 = p1.Next
        } else if p1.Val > p2.Val {
            p.Next = p2
            p2 = p2.Next
        }
        p = p.Next
    }
    if p1 != nil {
       p.Next =p1 
    }

    if p2 != nil {
        p.Next = p2
    }
        
    return dumy.Next
}

