package leetcode

// https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order/

import "sort"

type intSlice []int

func (s intSlice) Len() int {
        return len(s)
}

func (s intSlice) Less(i, j int) bool {
        return s[i] > s[j]
}

func (s intSlice) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
}

func MinSubsequence(nums []int) []int {

        var s intSlice = nums

        sort.Sort(s)

        sum := 0

        for idx := 0; idx < len(s); idx++ {
                sum += s[idx]
        }

        res := make([]int, 0)
        resSum := 0
        for idx := 0; idx < len(s); idx++ {

                res = append(res, s[idx])
                resSum += s[idx]
                if resSum > sum-resSum {
                        break
                }
        }

        return res
}
