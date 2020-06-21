package leetcode

// https://leetcode-cn.com/problems/count-number-of-teams/

func NumTeams(rating []int) int {
        n := len(rating)
        if n < 3 {
                return 0
        }
        num := 0
        for j := 1; j <= n-2; j++ {
                left, right := 0, 0
                for k := j + 1; k <= n-1; k++ {
                        if rating[k] > rating[j] {
                                right += 1
                        }
                }

                for i := j - 1; i >= 0; i-- {
                        if rating[i] < rating[j] {
                                left += 1
                        }
                }

                num += right*left + (j-left)*( n-j-1-right)
        }

        return num
}
