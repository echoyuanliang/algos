package pratice

func ReverseList(srcList *List) (*List) {

    if srcList == nil {
        return srcList
    }

    if srcList.HeadNode == nil || srcList.HeadNode.NextNode == nil {
        return srcList
    }

    priv := srcList.HeadNode
    cur := priv.NextNode
    priv.NextNode = nil
    for cur != nil {
        tmp := cur.NextNode
        cur.NextNode = priv
        priv = cur
        cur = tmp
    }

    srcList.HeadNode = priv

    return srcList
}

func AddTwoList(list1 *List, list2 *List) (*List) {

    if list1 == nil || list1.HeadNode == nil {
        return list2
    }

    if list2 == nil || list2.HeadNode == nil {
        return list1
    }

    reverseList1 := ReverseList(list1)
    reverseList2 := ReverseList(list2)

    node1 := reverseList1.HeadNode
    node2 := reverseList2.HeadNode
    carry := 0

    var nextNode *ListNode

    for node1 != nil || node2 != nil {
        sum := 0
        if node1 != nil && node2 != nil {
            //fmt.Printf("fuck fuck")
            sum = node1.Val + node2.Val
            node1 = node1.NextNode
            node2 = node2.NextNode
        } else if node1 != nil {
            //fmt.Printf("fuck1")
            sum = node1.Val
            node1 = node1.NextNode
        } else {
            //fmt.Printf("fuck2")
            sum = node2.Val
            node2 = node2.NextNode
        }

        sum += carry

        if sum >= 10 {
            carry = 1
        } else {
            carry = 0
        }

        sum = sum % 10
        if nextNode == nil {
            nextNode = new(ListNode)
            nextNode.Val = sum
        } else {
            curNode := new(ListNode)
            curNode.Val = sum
            curNode.NextNode = nextNode
            nextNode = curNode
        }
    }

    if carry != 0 {
        curNode := new(ListNode)
        curNode.Val = carry
        curNode.NextNode = nextNode
        nextNode = curNode
    }

    return &List{HeadNode: nextNode}
}
