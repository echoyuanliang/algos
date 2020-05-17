package leetcode

func LengthOfLongestSubstring(s string) int {

    seeMap := make([]int, 256)

    for idx := range seeMap {
        seeMap[idx] = -1
    }
    maxLen := 0
    curLen := 0
    start := 0
    for idx, c := range s {
        if seeMap[c] >= start {
            if curLen > maxLen {
                maxLen = curLen
            }

            curLen = idx - seeMap[c]
            start = seeMap[c] + 1
        } else {
            curLen += 1
        }

        seeMap[c] = idx
    }

    if curLen > maxLen {
        maxLen = curLen
    }
    return maxLen
}
