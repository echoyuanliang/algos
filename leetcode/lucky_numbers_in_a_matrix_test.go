package leetcode

import (
        "testing"
)

import (
        "github.com/stretchr/testify/assert"
)

func TestLuckyNumbers(t *testing.T) {
        tests := []struct {
                Input [][]int
                Want  []int
        }{
                {
                        Input: [][]int{
                                {3, 7, 8},
                                {9, 11, 13},
                                {15, 16, 17},
                        },
                        Want: []int{15},
                },
                {
                        Input: [][]int{
                                {1, 10, 4, 2},
                                {9, 3, 8, 7},
                                {15, 16, 17, 12},
                        },
                        Want: []int{12},
                },
                {
                        Input: [][]int{
                                {7, 8},
                                {1, 2},
                        },
                        Want: []int{7},
                },
        }

        for _, c := range tests {
                assert.ElementsMatch(t, c.Want, LuckyNumbers(c.Input))
        }
}
