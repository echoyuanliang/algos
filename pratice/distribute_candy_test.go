package pratice

import (
    "testing"
    "math/rand"
    "time"
    "fmt"
    "strings"
    "strconv"
)

func genRatings(n int) []int {
    ratings := make([]int, n)
    idx := 0

    for ; idx < n; idx++ {
        ratings[idx] = rand.Intn(100) + 1
    }

    return ratings
}

func printIntSlice(s []int) string {
    strSlice := make([]string, len(s))

    for idx, val := range s {
        strSlice[idx] = strconv.Itoa(val)
    }

    return strings.Join(strSlice, ",")
}

func TestGetMinCandyDistribute(t *testing.T) {
    testCnt := 0

    for ; testCnt < 10; testCnt++ {
        rand.Seed(time.Now().UnixNano())
        rateCnt := rand.Intn(10) + 1
        ratings := genRatings(rateCnt)
        candies := GetMinCandyDistribute(ratings)

        idx := 1

        for ; idx < rateCnt; idx++ {
            fail := false
            if ratings[idx] > ratings[ idx-1] && candies[idx] <= candies[idx-1] {
                fail = true
            } else if ratings[idx] < ratings[ idx-1] && candies[idx] >= candies[idx-1] {
                fail = true
            } else if ratings[idx] == ratings[ idx-1] && candies[idx] != candies[idx-1] {
                fail = true
            }

            if fail {
                t.Fatalf("idx %d, %d; rates: %d , %d; candies %d , %d; fatal",
                    idx, idx-1, ratings[idx], ratings[idx-1], candies[idx], candies[idx-1])
            }
        }

        fmt.Printf("ratings: %s\n", printIntSlice(ratings))
        fmt.Printf("candies: %s\n", printIntSlice(candies))
    }
}
