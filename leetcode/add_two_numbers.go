package leetcode


type ListNode struct{
    Val int
    NextNode *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{

    l1Cur, l2Cur := l1, l2
    resHead := &ListNode{0, nil}
    resCur := resHead

    for{
        resSum := resCur.Val
        if l1Cur != nil && l2Cur != nil{
            resSum += l1Cur.Val + l2Cur.Val
        }else if l1Cur != nil {
            resSum += l1Cur.Val
        }else if l2Cur != nil{
            resSum += l2Cur.Val
        }

        resCur.Val = resSum % 10

        nextNode :=   &ListNode{resSum / 10, nil}
        resCur.NextNode = nextNode

        if l1Cur != nil{
            l1Cur = l1Cur.NextNode
        }

        if l2Cur != nil{
            l2Cur = l2Cur.NextNode
        }

        if nextNode.Val == 0 && l1Cur == nil && l2Cur == nil{
            break
        }

        resCur = nextNode
    }

    return resHead
}