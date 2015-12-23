package main

import (
    "fmt"
    "github.com/baolidakai/hello/util"
)

// 232. Implement Queue using Stacks
type myque struct {
    stack1 stack.Stack
    stack2 stack.Stack
}

func (q *myque) Size() int {
    return q.stack1.Size() + q.stack2.Size()
}

func (q *myque) Push(x int) {
    q.stack1.Push(x)
}

func (q *myque) Pop() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop().(int))
        }
    }
    return q.stack2.Pop().(int)
}

func (q *myque) Peek() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop().(int))
        }
    }
    return q.stack2.Peek().(int)
}

func (q *myque) Empty() bool {
    return q.Size() == 0
}

func main() {
    queue := new(myque)
    // Test the queue (stack)
    queue.Push(2)
    queue.Push(3)
    queue.Push(4)
    queue.Pop()
    queue.Push(5)
    for queue.Size() != 0 {
        fmt.Println(queue.Pop())
    }
}
