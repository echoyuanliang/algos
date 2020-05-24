package leetcode

// https://leetcode-cn.com/problems/minimum-window-substring/

func MinWindow(s string, t string) string {
        hash := make([]int, 256)
        l, count, max, results := 0, len(t), len(s)+1, ""

        for i := 0; i < len(t); i++ {
                hash[t[i]]++
        }

        for r := 0; r < len(s); r++ {
                hash[s[r]]--
                if hash[s[r]] >= 0 {
                        count--
                }

                for l < r && hash[s[l]] < 0 {

                        hash[s[l]]++
                        l++

                }

                if count == 0 && r-l+1 < max {
                        max = r - l + 1
                        results = s[l:r+1]
                }
        }

        return results

}
