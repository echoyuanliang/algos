package sort

func qSort(sortArray []int, left, right int) {
    for left < right {
        key := sortArray[(left+right)/2]
        i := left
        j := right

        for {
            for sortArray[i] < key {
                i++
            }
            for sortArray[j] > key {
                j--
            }
            if i >= j {
                break
            }
            sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
        }

        qSort(sortArray, left, i-1)
        left = j+1
    }
}

func qSort2(src []int, first, last int) {
    flag := first
    left := first
    right := last

    if first >= last {
        return
    }

    for first < last {
        for first < last {
            if src[last] < src[flag]{
                src[last], src[flag] = src[flag], src[last]
                flag = last
                break
            }

            last -= 1
        }

        for first < last {
           if src[first] > src[flag]{
                src[first], src[flag] = src[flag], src[first]
                flag = first
                break
            }

            first += 1
        }
    }

    qSort2(src, left, flag-1)
    qSort2(src, flag+1, right)
}

func QSort(arr []int)  {
    qSort(arr, 0, len(arr) - 1)
}

func QSort2(arr []int){
    qSort2(arr, 0, len(arr) - 1)
}