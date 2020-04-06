package pratice

func CheckListHasCycle(l *List) (*ListNode) {
    if l == nil || l.HeadNode == nil || l.HeadNode.NextNode == nil {
        return nil
    }

    slow := l.HeadNode
    fast := l.HeadNode

    for {
        if fast == nil || fast.NextNode == nil {
            return nil
        }

        slow = slow.NextNode
        fast = fast.NextNode.NextNode

        if slow == fast {
            break
        }

    }

    fast = l.HeadNode

    for {
        if fast == slow {
            return fast
        }

        fast = fast.NextNode
        slow = slow.NextNode
    }
}

func CheckNoCycleListIntersect(l1 *List, l2 *List, endNode *ListNode) (*ListNode) {
    if l1 == nil || l1.HeadNode == nil {
        return nil
    }

    if l2 == nil || l2.HeadNode == nil {
        return nil
    }

    end1 := l1.HeadNode
    len1 := 1

    for end1.NextNode != endNode {
        end1 = end1.NextNode
        len1 += 1
    }

    end2 := l2.HeadNode
    len2 := 1

    for end2.NextNode != endNode {
        end2 = end2.NextNode
        len2 += 1
    }

    if end1 != end2 {
        return nil
    }

    cur1 := l1.HeadNode

    if len1 > len2 {
        fastStep := len1 - len2

        for fastStep != 0 {
            cur1 = cur1.NextNode
            fastStep -= 1
        }
    }

    cur2 := l2.HeadNode

    if len2 > len1 {
        fastStep := len2 - len1
        for fastStep != 0 {
            cur2 = cur2.NextNode
            fastStep -= 1
        }
    }

    for cur1 != cur2 {
        cur1 = cur1.NextNode
        cur2 = cur2.NextNode
    }

    return cur1
}

func CheckCycleListIntersect(l1 *List, l2 *List, loopStart1 *ListNode, loopStart2 *ListNode) (*ListNode) {

    if loopStart1 == loopStart2 {
        return CheckNoCycleListIntersect(l1, l2, loopStart1)
    }

    loop1Cur := loopStart1
    for loop1Cur != loopStart2 {
        loop1Cur = loop1Cur.NextNode

        if loop1Cur == loopStart1 {
            return nil
        }
    }

    return loop1Cur

}

func CheckListIntersect(l1 *List, l2 *List) (*ListNode) {
    if l1 == nil || l1.HeadNode == nil {
        return nil
    }

    if l2 == nil || l2.HeadNode == nil {
        return nil
    }

    loopStartL1 := CheckListHasCycle(l1)
    loopStartL2 := CheckListHasCycle(l2)

    if loopStartL1 == nil && loopStartL2 == nil {
        return CheckNoCycleListIntersect(l1, l2, nil)
    } else if loopStartL1 != nil && loopStartL2 != nil {
        return CheckCycleListIntersect(l1, l2, loopStartL1, loopStartL2)
    }

    return nil
}
