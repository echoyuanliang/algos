package sort

func MergeSort(arr []int){

    if len(arr) < 2{
        return
    }

    mid := len(arr) / 2
    MergeSort(arr[:mid])
    MergeSort(arr[mid:])

    tmpArr := make([]int, len(arr))
    start1, end1 := 0, mid
    start2, end2 := mid, len(arr)
    tmpIdx := 0

    for start1 < end1 && start2 < end2 {
       if arr[start1] > arr[start2]{
           tmpArr[tmpIdx] = arr[start2]
           start2++
       }else{
           tmpArr[tmpIdx] = arr[start1]
           start1++
       }

       tmpIdx++
    }

    for start1 < end1{
        tmpArr[tmpIdx] = arr[start1]
        tmpIdx++
        start1++
    }

    for start2 < end2{
        tmpArr[tmpIdx] = arr[start2]
        tmpIdx++
        start2++
    }

    copy(arr, tmpArr)
}