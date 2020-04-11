package tree

import (
    "container/list"
    "github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
    Value interface{}
    Left  *Node
    Right *Node
}

type Tree struct {
    Root *Node
}

func (tree *Tree) GetPreOrder() *list.List {
    stack := arraystack.Stack{}
    preOrderList := list.New()
    head := tree.Root

    if head != nil {
        stack.Push(head)
    }

    for ! stack.Empty() {

        value, _ := stack.Pop()
        head, _ = value.(*Node)
        preOrderList.PushBack(head.Value)

        if head.Right != nil {
            preOrderList.PushBack(head.Right)
        }

        if head.Left != nil {
            preOrderList.PushBack(head.Left)
        }
    }

    return preOrderList
}

func (tree *Tree) GetInOrder() *list.List {
    stack := arraystack.Stack{}
    inOrderList := list.New()
    head := tree.Root

    for ! stack.Empty() || head != nil {

        if head != nil {
            stack.Push(head)
            head = head.Left
        } else {
            value, _ := stack.Pop()
            head, _ = value.(*Node)
            inOrderList.PushBack(head.Value)
            head = head.Right
        }
    }

    return inOrderList
}

func (tree *Tree) GetPosOrder() *list.List {

    posOrderList := list.New()
    head := tree.Root

    if head != nil {
        stack := arraystack.Stack{}
        stack.Push(head)
        var cur *Node

        for ! stack.Empty() {
            value, _ := stack.Peek()
            cur, _ = value.(*Node)

            if cur.Left != nil && head != cur.Left && head != cur.Right {
                stack.Push(cur.Left)
            } else if cur.Right != nil && head != cur.Right {
                stack.Push(cur.Right)
            } else {
                stack.Pop()
                posOrderList.PushBack(cur.Value)
                head = cur
            }
        }

    }

    return posOrderList
}
