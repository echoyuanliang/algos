package sort


func InsertSort(arr []int){

    if len(arr) < 2{
        return
    }

    for endIdx :=1; endIdx < len(arr); endIdx++{
        temp := arr[endIdx]
        swapIdx := endIdx
        for ; swapIdx > 0 && arr[swapIdx-1] > temp; swapIdx--{
            arr[swapIdx] = arr[swapIdx-1]
        }

        arr[swapIdx] = temp
    }
}