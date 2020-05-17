package leetcode

// https://leetcode-cn.com/problems/course-schedule-ii/

import "container/list"

func FindOrder(numCourses int, prerequisites [][]int) []int {

    topRes := make([]int, 0)

    inDegree := make([]int, numCourses)

    l := list.New()

    dependencyMap := make(map[int][]int)

    for _, edge := range prerequisites {
        inDegree[edge[0]] += 1

        if _, ok := dependencyMap[edge[1]]; ok {
            dependencyMap[edge[1]] = append(dependencyMap[edge[1]], edge[0])
        } else {
            dependencyMap[edge[1]] = []int{edge[0]}
        }

    }

    for node, degree := range inDegree {
        if degree == 0 {
            l.PushFront(node)
        }
    }

    for l.Len() != 0 {
        e := l.Front()
        curVal := e.Value.(int)
        topRes = append(topRes, curVal)
        l.Remove(e)
        for _, dependency := range dependencyMap[curVal] {
            inDegree[dependency] -= 1

            if inDegree[dependency] == 0 {
                l.PushFront(dependency)
            }
        }

    }

    if len(topRes) < numCourses {
        return make([]int, 0)
    }
    return topRes
}
