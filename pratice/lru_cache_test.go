package pratice

import (
    "testing"
    "math/rand"
    "time"
)

func TestNewLruCache(t *testing.T) {
    capacity := 10
    visitRecord := make([]int, capacity*100)

    lruCache := NewLruCache(capacity)
    for idx := range visitRecord {
        rand.Seed(time.Now().UnixNano())
        visitRecord[idx] = rand.Intn(capacity * 2)
        lruCache.Put(visitRecord[idx], idx)

        func() {
            cacheItems := lruCache.GetAll()
            stop := idx - capacity + 1
            if stop < 0 {
                stop = 0
            }
            for recentIdx := idx; recentIdx >= stop; recentIdx-- {
                visitKey := visitRecord[recentIdx]

                if _, ok := cacheItems[visitKey]; ! ok {
                    t.Errorf("%d : %d -> %d error", idx, visitKey, recentIdx)
                }
            }
        }()

    }

}
