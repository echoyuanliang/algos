package leetcode

// https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/

func MaxIncreaseKeepingSkyline(grid [][]int) int {
        rowMax := make([]int, len(grid))
        colMax := make([]int, len(grid))

        for row := 0; row < len(grid); row++ {
                for col := 0; col < len(grid); col ++ {
                        if grid[row][col] > rowMax[row] {
                                rowMax[row] = grid[row][col]
                        }

                        if grid[row][col] > colMax[col] {
                                colMax[col] = grid[row][col]
                        }
                }
        }

        inc := 0

        for row := 0; row < len(grid); row++ {
                for col := 0; col < len(grid); col++ {
                        if grid[row][col] >= rowMax[row] || grid[row][col] >= colMax[col] {
                                continue
                        }

                        maxHigh := rowMax[row]

                        if maxHigh > colMax[col] {
                                maxHigh = colMax[col]
                        }

                        inc += maxHigh - grid[row][col]

                }

        }

        return inc
}
