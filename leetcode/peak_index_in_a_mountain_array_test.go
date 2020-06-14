package leetcode

import (
        "testing"
)

import (
        "github.com/stretchr/testify/assert"
)

func TestPeakIndexInMountainArray(t *testing.T) {
        tests := []struct {
                Input []int
                Want  int
        }{
                {
                        Input: []int{0, 1, 0},
                        Want:  1,
                },
                {
                        Input: []int{0, 2, 1, 0},
                        Want:  1,
                },
                {
                        Input: []int{0, 1, 2, 3, 4, 5, 6},
                        Want:  6,
                },
                {
                        Input: []int{7, 6, 5, 4, 3, 2, 1, 0},
                        Want:  0,
                },
                {
                        Input: []int{0, 6, 5, 4, 3},
                        Want:  1,
                },
                {
                        Input: []int{0, 1, 2, 4, 3, 2},
                        Want:  3,
                },
                {
                        Input: []int{40,48,61,75,100,99,98,39,30,10},
                        Want: 4,
                },
        }

        for _, c := range tests {
                PeakIndexInMountainArray(c.Input)

                assert.Equal(t, c.Want, PeakIndexInMountainArray(c.Input))
        }
}
