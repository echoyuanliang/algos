package leetcode

// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

func PeakIndexInMountainArray(A []int) int {
        length := len(A)

        if length == 1 {
                return A[0]
        }

        if length == 2 {
                if A[0] > A[1] {
                        return A[0]
                }

                return A[1]
        }

        idx := length / 2
        stop := length
        start := -1
        for idx > start && idx < stop {
                if idx > start+1 && A[idx] < A[idx-1] {
                        stop = idx
                        idx = (idx + start) / 2
                        continue
                } else if idx < stop-1 && A[idx] < A[idx+1] {
                        start = idx
                        idx = (stop + idx) / 2
                        continue
                } else {
                        break
                }
        }

        return idx
}
