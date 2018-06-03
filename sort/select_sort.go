package sort

func SelectSort(arr []int){
    rightIdx := len(arr)
    for leftIdx :=0; leftIdx<rightIdx; leftIdx++{
        minIdx := leftIdx
        for loopIdx := leftIdx+1; loopIdx<rightIdx; loopIdx++{
            if arr[loopIdx] < arr[minIdx]{
                minIdx = loopIdx
            }
        }

        if minIdx != leftIdx{
            arr[leftIdx], arr[minIdx] = arr[minIdx], arr[leftIdx]
        }
    }
}