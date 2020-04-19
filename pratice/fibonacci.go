package pratice

import (
    "math/big"
    "strconv"
)


func GetFibonacci(n uint64)  string{

    prev1 := big.NewInt(1)
    prev2 := big.NewInt(0)

    if n < 2{
        return strconv.Itoa(int(n))
    }

    var idx uint64 = 2
    for ; idx < n; idx++{
        prev2.Add(prev2, prev1)
        prev2, prev1 = prev1, prev2
    }

    prev2.Add(prev2, prev1)
    return prev2.String()
}