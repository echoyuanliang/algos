package leetcode

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.


Example:

    Given nums = [2, 7, 11, 15], target = 9,

    Because nums[0] + nums[1] = 2 + 7 = 9,
    return [0, 1].
*/

func HashTwoSum(nums []int, target int) ([]int){
    numIdxMap := make(map[int]int)

    for i:= 0; i < len(nums); i++{
        tupleValue := target - nums[i]

        if tupleIdx, ok := numIdxMap[tupleValue]; ok{
            return []int{i, tupleIdx}
        }

        numIdxMap[nums[i]] = i
    }

    return []int{-1, -1}
}


