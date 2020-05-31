package leetcode

import (
        "testing"
        "fmt"
)

func TestFindDiagonalOrder(t *testing.T) {

        type testCase struct {
                Input  [][]int
                Expect []int
        }

        var cases = []*testCase{
                {
                        Input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
                        Expect: []int{1, 4, 2, 7, 5, 3, 8, 6, 9},
                },
                {
                        Input:  [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}},
                        Expect: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16},
                },
                {
                        Input:  [][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}},
                        Expect: []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11},
                },
                {
                        Input:  [][]int{{1, 2, 3, 4, 5, 6}},
                        Expect: []int{1, 2, 3, 4, 5, 6},
                },
                {
                        Input:  [][]int{{5, 6, 3, 10}, {9}, {1, 19}, {9, 9, 2}},
                        Expect: []int{5, 9, 6, 1, 3, 9, 19, 10, 9, 2},
                },
                {
                        Input:  [][]int{{6}, {8}, {6, 1, 6, 16}},
                        Expect: []int{6, 8, 6, 1, 6, 16},
                },
        }

        for idx, c := range cases {
                output := FindDiagonalOrder(c.Input)
                fmt.Println(c.Expect)
                fmt.Println(output)
                if len(output) != len(c.Expect) {
                        t.Fatalf("fail at %d, len %d != %d", idx, len(output), len(c.Expect))
                }

                for j, o := range output {
                        if o != c.Expect[j] {
                                t.Fatalf("fail at %d, %d, %d != %d", idx, j, o, c.Expect[j])
                        }
                }

        }
}
