package pratice

import (
    "fmt"
    "container/list"
)

func permutation(arr []int, cursor int, end int) *list.List{

    permutationRes := list.New()

    if cursor == end{
        permutationRes.PushBack(fmt.Sprint(arr))
    }else{
        for i := cursor; i <= end; i++{
            arr[i], arr[cursor] = arr[cursor], arr[i]
            permutationRes.PushBackList(permutation(arr, cursor+1, end))
            arr[cursor], arr[i] = arr[i], arr[cursor]
        }
    }

    return permutationRes
}

func Permutation(arr []int) *list.List{
    return permutation(arr, 0, len(arr) - 1)
}
