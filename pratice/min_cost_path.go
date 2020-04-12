package pratice

import (
    "math"
    "github.com/toolkits/slice"
)


/**
    https://www.geeksforgeeks.org/min-cost-path-dp-6/?ref=lbp
*/

func GetMinCostPath(data [][]float64, m, n int) float64 {
    if m == 0 || n == 0{
        return 0
    }

    if m == 1{
        return slice.SumFloat64(data[0])
    }

    if n == 1{
        sum := 0.0
        for _, val := range data{
            sum += val[0]
        }

        return sum
    }


    helpData := make([]float64, m)
    helpData[0] = data[0][0]
    for i := 1; i < m; i++ {
        helpData[i] = helpData[i-1] + data[i][0]
    }

    for j := 1; j < n; j++ {
        helpData[0] = helpData[0] + data[0][j]
       for i := 1; i < m; i++{
           helpData[i] = math.Min(helpData[i-1], helpData[i]) + data[i][j]
       }

    }

    return helpData[m-1]
}
