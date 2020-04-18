package pratice

/**
    https://www.interviewbit.com/problems/distribute-candy/
 */

func GetMinCandyDistribute(ratings []int) []int {

    if len(ratings) < 2 {
        return []int{len(ratings)}
    }

    candies := make([]int, len(ratings))
    candies[0] = 1
    idx := 1
    for ; idx < len(ratings); idx++ {
        if ratings[idx] > ratings[idx-1] {
            candies[idx] = candies[idx-1] + 1
        } else if ratings[idx] == ratings[idx-1] {
            candies[idx] = candies[idx-1]
        } else {
            candies[idx] = 1
        }
    }

    idx = len(ratings) - 2

    for ; idx >= 0; idx-- {
        if ratings[idx] > ratings[idx+1] {
            candies[idx] = func(l int, r int) int {
                if l > r {
                    return l
                }

                return r
            }(candies[idx], candies[idx+1]+1)
        } else if ratings[idx] == ratings[idx+1] {
            candies[idx] = candies[idx+1]
        }
    }

    return candies

}
