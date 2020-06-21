package leetcode

import (
        "testing"
)

import (
        "github.com/stretchr/testify/assert"
)

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
        tests := []struct {
                Input [][]int
                Want  int
        }{
                {
                        Input: [][]int{
                                {3, 0, 8, 4},
                                {2, 4, 5, 7},
                                {9, 2, 6, 3},
                                {0, 3, 1, 0},
                        },
                        Want: 35,
                },
        }

        for _, c := range tests {

                assert.Equal(t, c.Want, MaxIncreaseKeepingSkyline(c.Input))
        }
}
