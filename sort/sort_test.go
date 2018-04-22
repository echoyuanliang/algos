package sort

import (
    "testing"
    "math/rand"
)

func TestQSort(t *testing.T) {
    randIntegers := rand.Perm(1000000)
    QSort(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "must be ", val)
        }
    }

}


func TestQSort2(t *testing.T) {
    randIntegers := rand.Perm(1000000)
    QSort2(randIntegers)

    for idx, val := range randIntegers{
        if idx != val{
            t.Error("idx ", idx, "must be ", val)
        }
    }

}
