package pratice

import (
    "math/rand"
    "time"
)

func ReservoirSampling(k int, pool []int) []int {
    pLen := len(pool)

    sample := make([]int, k)
    for i := 0; i < pLen; i++ {
        if i < k {
            sample[i] = pool[i]
        } else {
            rand.Seed(time.Now().UnixNano())
            choiceRate := rand.Intn(i)
            if choiceRate < k {
                dropIdx := rand.Intn(k)
                sample[dropIdx] = pool[i]
            }
        }
    }

    return sample
}
