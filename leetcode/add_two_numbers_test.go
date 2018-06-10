package leetcode

import (
    "testing"
    "math/rand"
    "time"
    "sync"
)

const TestMax = 5000000

func genListByNumbers(num int) *ListNode {
    if num < 0 {
        return nil
    }

    head := &ListNode{num % 10, nil}
    cur := head
    num /= 10
    posNum := 1
    for num > 0{
        nextNode := &ListNode{num % 10, nil}
        cur.NextNode = nextNode
        cur = cur.NextNode
        num /= 10
        posNum += 1
    }

    return head
}


func listToNumber(l *ListNode) int{
    if l == nil{
        return 0
    }

    res := l.Val
    curWeight := 1
    curPos := l.NextNode

    for curPos != nil{
        res += curPos.Val * curWeight * 10
        curWeight *= 10
        curPos = curPos.NextNode
    }

    return res
}


func testAddTwoNumbers(t *testing.T){
    rand.Seed(time.Now().UnixNano())
    rNum, lNum := rand.Intn(TestMax), rand.Intn(TestMax)
    l1 := genListByNumbers(rNum)
    l2 := genListByNumbers(lNum)

    resList := AddTwoNumbers(l1, l2)
    resNum := listToNumber(resList)
    if resNum != rNum + lNum{
        t.Fatal(resNum, "!=", rNum, "+", lNum)
    }
}


func TestAddTwoNumbers(t *testing.T) {
    const TestCnt = 100
    resLock := sync.WaitGroup{}
    for cnt := 0; cnt < TestCnt; cnt++ {
        go func() {
            resLock.Add(1)
            testAddTwoNumbers(t)
            resLock.Done()
        }()
    }

    resLock.Wait()
}