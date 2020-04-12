package pratice

import "math"

func GetCoinChangeWays(target int, coinValues []int) int {

    targetList := make([]int, target+1)
    targetList[0] = 1

    for t := 1; t <= target; t++ {

        if t%coinValues[0] == 0 {
            targetList[t] = 1
        } else {
            targetList[t] = 0
        }
    }

    for c := 1; c < len(coinValues); c++ {
        for t := 1; t <= target; t++ {
            if t >= coinValues[c] {
                targetList[t] = targetList[t] + targetList[t-coinValues[c]]
            }
        }
    }

    return targetList[target]
}

func GetMiniNumCoins(target int, coinValues []int) int {

    targetList := make([]int, target+1)
    targetList[0] = 0

    max := math.MaxInt64 - 10
    for t := 1; t <= target; t++ {

        if t%coinValues[0] == 0 {
            targetList[t] = t / coinValues[0]
        } else {
            targetList[t] = max
        }
    }

    for c := 1; c < len(coinValues); c++ {
        for t := 1; t <= target; t++ {
            if t >= coinValues[c] && targetList[t-coinValues[c]]+1 < targetList[t] {
                targetList[t] = targetList[t-coinValues[c]] + 1
            }
        }
    }

    return targetList[target]
}
