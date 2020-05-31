package leetcode

// https://leetcode-cn.com/problems/diagonal-traverse-ii/submissions/

func FindDiagonalOrder(nums [][]int) []int {
        numsLen := len(nums)

        if numsLen == 0 {
                return nil
        }

        if numsLen == 1 {
                return nums[0]
        }

        numsMap := make(map[int][]int)

        for i := 0; i < numsLen; i++ {
                for j := 0; j < len(nums[i]); j++ {
                        numsMap[i+j] = append(numsMap[i+j], nums[i][j])
                }
        }

        res := make([]int, 0, len(numsMap))
        for i := 0; i < len(numsMap); i++ {
                for j := len(numsMap[i]) - 1; j >= 0; j-- {
                        res = append(res, numsMap[i][j])
                }
        }

        return res
}
