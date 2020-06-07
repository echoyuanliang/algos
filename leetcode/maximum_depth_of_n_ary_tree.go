package leetcode

import "container/list"

// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/

type Node struct {
        Val      interface{}
        Children []*Node
}

func MaxDepthDfs(root *Node) int {
        if root == nil {
                return 0
        }

        depth := 0
        for _, node := range root.Children {
                d := MaxDepthDfs(node)

                if d > depth {
                        depth = d
                }
        }

        return depth + 1
}

func MaxDepthBfs(root *Node) int {
        depth := 0
        if root == nil {
                return depth
        }

        l := list.New()

        l.PushBack(root)
        for l.Len() != 0 {
                count := l.Len()
                depth += 1
                for count > 0 {
                        node := l.Remove(l.Front()).(*Node)
                        for _, c := range node.Children {
                                l.PushBack(c)
                        }
                        count --
                }

        }

        return depth
}
