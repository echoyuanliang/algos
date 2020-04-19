package pratice

import (
    "time"
    "testing"
    "math/rand"
    "fmt"
)

func TestGetFibonacci(t *testing.T) {

    for testCnt := 0; testCnt < 10; testCnt++ {
        rand.Seed(time.Now().UnixNano())
        fibInput := rand.Intn(100) + 1
        fmt.Printf("round %d, fibInput %d:\n", testCnt, fibInput)
        fibResult := GetFibonacci(uint64(fibInput))
        fmt.Printf("fibResult: %s\n", fibResult)
    }
}
