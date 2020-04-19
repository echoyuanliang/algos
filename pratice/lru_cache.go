package pratice

import (
    "container/list"
    "errors"
)

type LruCache struct {
    capacity  int
    container map[int]int
    keyList   *list.List
    keyMap    map[int]*list.Element
}

func NewLruCache(capacity int) *LruCache {
    cache := &LruCache{
        capacity:  capacity,
        container: make(map[int]int),
        keyList:   list.New(),
        keyMap:    make(map[int]*list.Element),
    }
    return cache
}

func (lruCache *LruCache) Put(key int, value int) {

    if _, ok := lruCache.container[key]; ok {
        lruCache.keyList.MoveToFront(lruCache.keyMap[key])
    } else {
        lruCache.keyList.PushFront(key)
    }
    lruCache.container[key] = value
    lruCache.keyMap[key] = lruCache.keyList.Front()

    if len(lruCache.container) > lruCache.capacity {
        dropNode := lruCache.keyList.Back()
        dropKey := dropNode.Value.(int)
        delete(lruCache.container, dropKey)
        lruCache.keyList.Remove(dropNode)
        delete(lruCache.keyMap, dropKey)
    }
}

func (lruCache *LruCache) Get(key int) (error, int) {
    if val, ok := lruCache.container[key]; ok {
        node := lruCache.keyMap[key]
        lruCache.keyList.MoveToFront(node)
        return nil, val
    }

    return errors.New("not exists"), -1
}

func (lruCache *LruCache) GetAll() map[int]int {

    cacheAll := make(map[int]int, len(lruCache.container))
    for key, value := range lruCache.container {
        cacheAll[key] = value
    }

    return cacheAll
}
