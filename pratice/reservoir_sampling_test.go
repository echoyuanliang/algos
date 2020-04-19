package pratice

import (
    "testing"
    "strconv"
    "strings"
    "fmt"
)

func TestReservoirSampling(t *testing.T) {

    n := 1000

    pool := make([]int, n)
    for idx := 0; idx < n; idx++ {
        pool[idx] = idx
    }

    for k := 10; k <= 100; k += 10 {
        sample := ReservoirSampling(k, pool)

        func(s []int) {
            strSlice := make([]string, len(s))

            for idx, val := range s {
                strSlice[idx] = strconv.Itoa(val)
            }

            fmt.Println(strings.Join(strSlice, ","))
        }(sample)
    }
}
