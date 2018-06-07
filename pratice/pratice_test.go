package pratice

import (
    "testing"
    "math/rand"
    "container/list"
    "errors"
    "fmt"
)

const TestCnt = 10

func factorial(x int) int {
    if x <= 2 {
        return x
    }

    return x * factorial(x-1)
}

func validatePermRes(res *list.List, expectedLen int) (error) {
    if res.Len() != expectedLen {
        return errors.New(fmt.Sprintf("permutation of %d must be %d", res.Len(), expectedLen))
    }

    permMap := make(map[string]interface{})

    for permItem := res.Front(); permItem != nil; permItem = permItem.Next() {
        permMap[permItem.Value.(string)] = nil
    }

    if res.Len() != len(permMap) {
        return errors.New(fmt.Sprintf("dumplicate permutation gen"))
    }

    return nil
}

func TestRecursePermutation(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)

    expectedLen := factorial(TestCnt)

    err := validatePermRes(RecursePermutation(randIntegers), expectedLen)

    if err != nil {
        t.Fatal(err)
    }
}


func TestLoopPermutation(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)

    expectedLen := factorial(TestCnt)

    err := validatePermRes(LoopPermutation(randIntegers), expectedLen)

    if err != nil {
        t.Fatal(err)
    }
}

