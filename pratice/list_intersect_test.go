package pratice

import (
    "testing"
    "fmt"
)

func TestNullList(t *testing.T) {
    fmt.Println("TestNullList start")
    l := new(List)
    res := CheckListIntersect(nil, l)

    if res != nil {
        t.Fatal("check null list failed")
    }

    res = CheckListIntersect(l, l)

    if res != nil {
        t.Fatal("check null list failed")
    }
}

func TestHeadList(t *testing.T) {
    fmt.Println("TestHeadList start")
    l1 := &List{HeadNode: &ListNode{Val: 5}}
    l2 := &List{HeadNode: &ListNode{Val: 3}}

    res := CheckListIntersect(l1, l2)

    if res != nil {
        t.Fatal("check only head no cycle list no intersect failed")
    } else {
        t.Log("check only head no cycle list no intersect ok")
    }

    l1.HeadNode.NextNode = l1.HeadNode

    res = CheckListIntersect(l1, l2)

    if res != nil {
        t.Fatal("check only head cycle and no cycle list no intersect failed")
    } else {
        t.Log("check only head cycle and no cycle list no intersect ok")
    }

    l2.HeadNode.NextNode = l2.HeadNode
    res = CheckListIntersect(l1, l2)
    if res != nil {
        t.Fatal("check only head cycle list no intersect failed")
    } else {
        t.Log("check only head cycle list no intersect ok")
    }

    l1.HeadNode = l2.HeadNode

    res = CheckListIntersect(l1, l2)
    if res != l1.HeadNode {
        t.Fatal("check only head cycle list intersect failed")
    } else {
        t.Log("check only head cycle list intersect ok")
    }

    l1.HeadNode.NextNode = nil
    res = CheckListIntersect(l1, l2)
    if res != l1.HeadNode {
        t.Fatal("check only head no cycle list intersect failed")
    } else {
        t.Log("check only head no cycle list intersect ok")
    }
}

func TestNormalList(t *testing.T) {
    fmt.Println("TestNormalList start")
    initVal := 0
    l1 := &List{HeadNode: &ListNode{Val: initVal, NextNode: nil}}
    n1 := l1.HeadNode

    for initVal != 5 {
        initVal += 1
        n1.NextNode = &ListNode{Val: initVal, NextNode: nil}
        n1 = n1.NextNode
    }

    initVal = 0
    l2 := &List{HeadNode: &ListNode{Val: initVal, NextNode: nil}}
    n2 := l2.HeadNode

    for initVal != 10 {
        initVal += 1
        n2.NextNode = &ListNode{Val: initVal, NextNode: nil}
        n2 = n2.NextNode
    }
    fmt.Println("check normal start")
    res := CheckListIntersect(l1, l2)
    if res != nil {
        t.Fatal("check no cycle list no intersect failed")
    }

    fmt.Println("check normal 1 end")
    n1.NextNode = l2.HeadNode
    res = CheckListIntersect(l1, l2)

    if res != n1.NextNode {
        t.Fatalf("check no cycle list intersect failed")
    }

    n1.NextNode = n2

    res = CheckListIntersect(l1, l2)

    if res != n1.NextNode {
        t.Fatalf("check no cycle list intersect failed")
    }

    n1.NextNode = l1.HeadNode
    n2.NextNode = l2.HeadNode

    res = CheckListIntersect(l1, l2)

    if res != nil {
        t.Fatalf("check  cycle list no intersect failed")
    }

    n1.NextNode = n2

    res = CheckListIntersect(l1, l2)
    if res != n2 && res != l2.HeadNode {
        t.Fatalf("check  cycle list intersect failed")
    }

}
