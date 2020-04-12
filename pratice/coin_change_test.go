package pratice

import (
    "testing"
    "fmt"
)

func TestCoinChange(t *testing.T) {
    coinValues := []int{1, 2, 5, 10}

    for target := 13; target < 194; target += 7 {
        n := GetMiniNumCoins(target, coinValues)
        w := GetCoinChangeWays(target, coinValues)
        fmt.Printf("target %d at least need %d coins, all %d ways\n", target, n, w)
    }

}
