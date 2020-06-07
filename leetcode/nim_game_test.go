package leetcode

import "testing"

func TestCanWinNim(t *testing.T) {

        type testCase struct {
                Input  int
                Expect bool
        }

        cases := []*testCase{
                {
                        4,
                        false,
                },
                {
                        5,
                        true,
                }, {
                        7,
                        true,
                },
        }

        for idx, c := range cases {
                if CanWinNim(c.Input) != c.Expect {
                        t.Errorf("%d input %d fail !", idx, c.Input)
                }

        }
}
