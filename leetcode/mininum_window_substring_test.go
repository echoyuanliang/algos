package leetcode

import "testing"

func TestMinWindow(t *testing.T) {
        type input struct {
                S string
                T string
        }

        type testCase struct {
                Input  *input
                Expect string
        }

        cases := []*testCase{
                {
                        &input{
                                S: "ADOBECODEBANC",
                                T: "ABC",
                        },
                        "BANC",
                },
                {
                        &input{
                                S: "AAAAAAAAAADOBECODEBANC",
                                T: "ABC",
                        },
                        "BANC",
                },
                {
                        &input{
                                S: "aa",
                                T: "aa",
                        },
                        "aa",
                },
                {
                        &input{
                                S: "a",
                                T: "aa",
                        },
                        "",
                },
                {
                        &input{
                                S: "aaaaaaaaaaaabbbbbcdd",
                                T: "abcdd",
                        },
                        "abbbbbcdd",
                },
        }

        for _, cs := range cases {
                output := MinWindow(cs.Input.S, cs.Input.T)

                if output != cs.Expect {
                        t.Errorf("input %s, %s, expect: %s, output: %s", cs.Input.S, cs.Input.T, cs.Expect, output)
                }
        }

}
