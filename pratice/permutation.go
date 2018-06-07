package pratice

import (
    "fmt"
    "container/list"
    "sort"
)

func recursePermutation(arr []int, cursor int, end int) *list.List{

    permutationRes := list.New()

    if cursor == end{
        permutationRes.PushBack(fmt.Sprint(arr))
    }else{
        for i := cursor; i <= end; i++{
            arr[i], arr[cursor] = arr[cursor], arr[i]
            permutationRes.PushBackList(recursePermutation(arr, cursor+1, end))
            arr[cursor], arr[i] = arr[i], arr[cursor]
        }
    }

    return permutationRes
}


func RecursePermutation(arr []int) *list.List{
    return recursePermutation(arr, 0, len(arr) - 1)
}


func reverse(arr []int){
    i := 0
    j := len(arr) -1

    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}


func LoopPermutation(arr []int) *list.List{

    sort.Ints(arr)
    arrLen := len(arr)
    permutationRes :=list.New()

    for{
        permutationRes.PushBack(fmt.Sprint(arr))
        i := arrLen - 2
        for ; i >= 0; i--{
            if arr[i] < arr[i+1] {
                break
            }else if i == 0{
                return permutationRes
            }
        }

        j := arrLen - 1
        for ; j > i; j--{
            if arr[j] > arr[i]{
                break
            }
        }

        arr[i], arr[j] = arr[j], arr[i]
        reverse(arr[i+1:])
    }

}