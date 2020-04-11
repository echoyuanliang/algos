package list

import (
    "testing"
    "math/rand"
    "time"
    "fmt"
)

const TestMax = 999999999999999999

func genListByNumbers(num int) *List {

    nextNode := &ListNode{num % 10, nil}
    num = num / 10
    for num > 0 {
        curNode := &ListNode{num % 10, nil}
        curNode.NextNode = nextNode
        nextNode = curNode
        num = num / 10
    }
    resList := new(List)
    resList.HeadNode = nextNode
    return resList
}

func listToNumber(l *List) int {
    if l == nil || l.HeadNode == nil {
        return 0
    }

    res := 0
    node := l.HeadNode
    for node != nil {
        res = res*10 + node.Val
        node = node.NextNode
    }

    return res
}

func testAddTwoList(t *testing.T, rNum int, lNum int) {

    l1 := genListByNumbers(rNum)
    l2 := genListByNumbers(lNum)
    resList := AddTwoList(l1, l2)
    resNum := listToNumber(resList)

    if rNum < 0 || lNum < 0 {
        return
    }
    if resNum != rNum+lNum {
        t.Fatal(resNum, "!=", rNum, "+", lNum)
    }
}

func TestAddTwoList(t *testing.T) {
    fmt.Print("TestAddTwoList start")
    const TestCnt = 1000
    testAddTwoList(t, 0, 9999)
    testAddTwoList(t, 9999, 0)
    for cnt := 0; cnt < TestCnt; cnt++ {
        rand.Seed(time.Now().UnixNano())
        rNum, lNum := rand.Intn(TestMax), rand.Intn(TestMax)
        testAddTwoList(t, rNum, lNum)
    }

}
