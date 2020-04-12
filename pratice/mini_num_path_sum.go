package pratice

import (
    "math"
    "fmt"
    "strings"
    "github.com/toolkits/slice"
)

func helpTest(arr []float64, rows int){
    s := make([]string, len(arr))
    for idx, val := range arr{
        s[idx] = fmt.Sprintf("%f", val)
    }

    fmt.Printf("rows %d:\n", rows)
    fmt.Println(strings.Join(s, " , "))
}


func GetMiniNumPathSum(data [][]float64, m, n int) float64 {
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
