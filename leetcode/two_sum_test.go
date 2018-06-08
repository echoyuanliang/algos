package leetcode

import (
    "testing"
    "math/rand"
    "fmt"
    "time"
    "sync"
)

const TestCnt = 50000

func testTwoSum(t *testing.T){
    rand.Seed(time.Now().UnixNano())
    randIntegers := rand.Perm(TestCnt)
    rIdx, lIdx := rand.Intn(TestCnt), rand.Intn(TestCnt)
    sum := randIntegers[rIdx] + randIntegers[lIdx]
    ret := HashTwoSum(randIntegers, sum)
    if randIntegers[ret[0]] + randIntegers[ret[1]] != sum{
        t.Fatal(fmt.Sprintf("nums %d:%d + nums %d:%d != sum %d",
            ret[0], randIntegers[ret[0]], ret[1], randIntegers[ret[1]], sum))
    }
}

func TestTwoSum(t *testing.T) {
    resLock := sync.WaitGroup{}

    for cnt := 0; cnt < TestCnt; cnt ++{
        go func(myCnt int){
            resLock.Add(1)
            testTwoSum(t)
            resLock.Done()
        }(cnt)

    }

    resLock.Wait()
}