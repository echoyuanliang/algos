package leetcode

import "testing"

var flagTests = []struct {
    in  string
    out int
}{
    {"''", 1},
    {"abcdba", 4},
    {"abcabcbb", 3},
    {"bbbbb", 1},
    {"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {

    for _, tt := range flagTests {
        res := LengthOfLongestSubstring(tt.in)

        if res != tt.out {
            t.Errorf("in %s, want %d, out %d", tt.in, tt.out, res)
        }
    }
}
