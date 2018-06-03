package pratice

import (
    "testing"
    "math/rand"
)


const TestCnt = 10

func factorial(x int) int{
    if x <= 2{
        return x
    }

    return x * factorial(x -1)
}

func TestPermutation(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)

    expectedLen := factorial(len(randIntegers))

    permRes := Permutation(randIntegers)
    if permRes.Len() != expectedLen{
        t.Fatal("permutation of", permRes.Len(), "must be ", expectedLen)
    }

    permMap := make(map[string]interface{})

    for permItem := permRes.Front(); permItem != nil; permItem = permItem.Next(){
        permMap[permItem.Value.(string)] = nil
    }

    if permRes.Len() != len(permMap){
        t.Fatal("permutation is not unique")
    }
}
