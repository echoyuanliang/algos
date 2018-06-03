package sort

func BubbleSort(arr []int)  {
    if len(arr) < 2 {
        return
    }


    toIdx := len(arr) - 1

    for fromIdx := 0; fromIdx < toIdx; fromIdx++{
        for tmpIdx := 0; tmpIdx < toIdx-fromIdx; tmpIdx++{
            if arr[tmpIdx] > arr[tmpIdx+1]{
                arr[tmpIdx], arr[tmpIdx+1] = arr[tmpIdx+1], arr[tmpIdx]
            }
        }
    }
}
