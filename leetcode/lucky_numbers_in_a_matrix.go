package leetcode

// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/

func LuckyNumbers(matrix [][]int) []int {

        m := len(matrix)
        n := len(matrix[0])
        var res []int

        for _, row := range matrix {
                min := 0
                for c := 1; c < n; c ++ {
                        if row[c] < row[min] {
                                min = c
                        }
                }

                flag := true

                for j := 0; j < m; j++ {
                        if matrix[j][min] > row[min] {
                                flag = false
                                break
                        }
                }

                if flag {
                        res = append(res, row[min])
                }

        }

        return res

}
