package sort

import (
    "testing"
    "math/rand"
)

const TestCnt = 100000

func TestQSort(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    QSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "can't be ", val)
        }
    }

}


func TestQSort2(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    QSort2(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "can't be ", val)
        }
    }

}

func TestBubbleSort(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    BubbleSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "cant't be ", val)
        }
    }
}

func TestMergeSort(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    MergeSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "cant't be ", val)
        }
    }
}

func TestSelectSort(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    SelectSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "cant't be ", val)
        }
    }
}

func TestInsertSort(t *testing.T) {
    randIntegers := rand.Perm(TestCnt)
    InsertSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "cant't be ", val)
        }
    }
}