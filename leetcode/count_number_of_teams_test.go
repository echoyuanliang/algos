package leetcode

import (
        "testing"
)

import (
        "github.com/stretchr/testify/assert"
)

func TestNumTeams(t *testing.T) {
        tests := []struct {
                Input []int
                Want  int
        }{
                {
                        Input: []int{2, 5, 3, 4, 1},
                        Want:  3,
                },
                {
                        Input: []int{2, 1, 3},
                        Want:  0,
                },
                {
                        Input: []int{1, 2, 3, 4},
                        Want:  4,
                },
        }

        for _, c := range tests {

                assert.Equal(t, c.Want, NumTeams(c.Input))
        }
}
