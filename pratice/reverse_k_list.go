package pratice

func ReverseKList(l *List, k int) (*List) {

    if l == nil || l.HeadNode == nil || l.HeadNode.NextNode == nil || k < 2 {
        return l
    }

    start := l.HeadNode
    end := start
    loopCnt := 1
    var lastTail *ListNode

    for end != nil {
        if loopCnt%k == 0 {
            cur := start
            next := cur.NextNode
            for cur != end {
                tmp := next.NextNode
                next.NextNode = cur
                cur = next
                next = tmp
            }

            if lastTail != nil {
                lastTail.NextNode = cur
            } else {
                l.HeadNode = cur
            }

            lastTail = start
            start = next
            end = next
        } else {
            end = end.NextNode
        }

        loopCnt += 1
    }

    if lastTail != nil {
        if start != end {
            lastTail.NextNode = start
        } else {
            lastTail.NextNode = nil
        }
    }

    return l
}
