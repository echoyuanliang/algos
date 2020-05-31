package leetcode

import (
        "testing"
)

func TestMinSubsequence(t *testing.T) {
        type testCase struct {
                Input  []int
                Expect []int
        }

        var cases = []*testCase{
                {
                        Input:  []int{4, 3, 10, 9, 8},
                        Expect: []int{2, 19},
                },
                {
                        Input:  []int{4, 4, 7, 6, 7},
                        Expect: []int{3, 20},
                },
                {
                        Input:  []int{6},
                        Expect: []int{1, 6},
                },
        }

        for idx, c := range cases {
                output := MinSubsequence(c.Input)

                if len(output) != c.Expect[0] {
                        t.Errorf("fail at %d, len %d", idx, len(output))
                }

                sum := 0
                for _, item := range output {
                        sum += item
                }

                if sum != c.Expect[1] {
                        t.Errorf("fail at %d, sum %d", idx, sum)
                }
        }
}
