package pratice

import (
    "testing"
    "fmt"
    "strconv"
    "strings"
)

func printList(l *List) (string) {

    var lStr []string
    if l != nil || l.HeadNode != nil {
        ptr := l.HeadNode

        for ptr != nil {
            lStr = append(lStr, strconv.Itoa(ptr.Val))
            ptr = ptr.NextNode
        }
    }

    return strings.Join(lStr, ",")
}

func copyList(l *List) (*List) {
    if l == nil {
        return l
    }

    cpList := &List{}

    ptr := l.HeadNode
    if ptr == nil {
        return cpList
    }

    cpList.HeadNode = &ListNode{Val: ptr.Val}
    cpPtr := cpList.HeadNode
    for ptr.NextNode != nil {
        cpPtr.NextNode = &ListNode{Val: ptr.NextNode.Val}
        ptr = ptr.NextNode
        cpPtr = cpPtr.NextNode
    }

    return cpList
}

func testHelper(l *List, len, k int) {
    cpList := copyList(l)
    oldListStr := printList(cpList)
    cpList = ReverseKList(cpList, k)
    newListStr := printList(cpList)
    fmt.Printf("%d,%d : %s ; %s\n", len, k, oldListStr, newListStr)
}

func TestReverseKList(t *testing.T) {
    fmt.Println("TestReverseKList start")
    initVal := 1
    l := &List{HeadNode: &ListNode{Val: initVal, NextNode: nil}}
    n := l.HeadNode
    testHelper(l, 1, 3)

    for initVal != 10 {
        initVal += 1
        n.NextNode = &ListNode{Val: initVal, NextNode: nil}
        n = n.NextNode

        testHelper(l, initVal, 10)
        testHelper(l, initVal, 5)
        testHelper(l, initVal, 4)
        testHelper(l, initVal, 3)
        testHelper(l, initVal, 2)
    }

}
