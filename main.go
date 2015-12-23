package main

import (
    "fmt"
    "github.com/baolidakai/hello/stack"
)

// 232. Implement Queue using Stacks
type myque struct {
    stack1 stack.Stack
}

func (q *myque) Size() int {
    return q.stack1.Size()
}

func (q *myque) Push(x int) {
    q.stack1.Push(x)
}

func (q *myque) Pop() int {
    return q.stack1.Pop().(int)
}

func (q *myque) Peek() int {
    rtn := q.stack1.Pop().(int)
    q.stack1.Push(rtn)
    return rtn
}

func (q *myque) Empty() bool {
    return q.Size() == 0
}

func main() {
    fmt.Println("tmp")
    queue := new(myque)
    // Test the queue (stack)
    fmt.Println(queue.Size())
    queue.Push(2)
    queue.Push(3)
    queue.Push(4)
    for queue.Size() != 0 {
        fmt.Println(queue.Pop())
    }
}
