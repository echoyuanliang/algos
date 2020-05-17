package leetcode

import (
    "testing"
)

func TestFindOrder(t *testing.T) {
    var input = [][]int{
        {1, 0},
    }

    output := FindOrder(2, input)
    t.Log(output)

    input = [][]int{
        {1, 0},
        {2, 0},
        {3, 1},
        {3, 2},
    }

    output = FindOrder(4, input)
    t.Log(output)

    input = [][]int{
        {1, 0},
        {0, 2},
        {2, 1},
    }

    output = FindOrder(3, input)
    t.Log(output)
}
