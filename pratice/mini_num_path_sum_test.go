package pratice

import (
    "testing"
    "math/rand"
    "fmt"
    "strings"
    "time"
)

func getTestArr(m, n int) [][]float64 {
    testArr := make([][]float64, m)

    for i := 0; i < m; i++ {
        testArr[i] = make([]float64, n)
        rand.Seed(time.Now().UnixNano())
        for j := 0; j < n; j++ {
            testArr[i][j] = float64(rand.Int() % 10)
        }
    }

    return testArr
}

func reverseTestArr(arr [][] float64, m, n int) {
    for i := 0; i < m-i-1; i++ {
        ir := m - i - 1
        for j := 0; j < n-j-1; j++ {
            jr := n - j - 1
            arr[i][j], arr[ir][jr] = arr[ir][jr], arr[i][j]
        }
    }
}

func printTestArr(arr [][] float64) {
    for _, v := range arr {
        l := make([]string, len(v))
        for idx, f := range v {
            l[idx] = fmt.Sprintf("%f", f)
        }
        fmt.Println(strings.Join(l, " , "))
    }
}

func TestGetMiniNumPathSum(t *testing.T) {
    m, n := 4, 5
    testArr := getTestArr(m, n)
    printTestArr(testArr)
    miniNumPathSum := GetMiniNumPathSum(testArr, m, n)
    fmt.Println(miniNumPathSum)
    reverseTestArr(testArr, m, n)
    printTestArr(testArr)
    miniNumPathSum = GetMiniNumPathSum(testArr, m, n)
    fmt.Println(miniNumPathSum)
}
