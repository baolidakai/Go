package queue

import "github.com/baolidakai/hello/stack"

//queue := new(queue.Queue)
//// Test the queue (stack)
//queue.Push(2)
//queue.Push(3)
//queue.Push(4)
//queue.Pop()
//queue.Push(5)
//for !queue.Empty() {
//	fmt.Println(queue.Pop())
//}

// 232. Implement Queue using Stacks
type Queue struct {
    stack1 stack.Stack
    stack2 stack.Stack
}

func (q *Queue) Size() int {
    return q.stack1.Size() + q.stack2.Size()
}

func (q *Queue) Push(x int) {
    q.stack1.Push(x)
}

func (q *Queue) Pop() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop().(int))
        }
    }
    return q.stack2.Pop().(int)
}

func (q *Queue) Peek() int {
    if q.stack2.IsEmpty() {
        for !q.stack1.IsEmpty() {
            q.stack2.Push(q.stack1.Pop().(int))
        }
    }
    return q.stack2.Peek().(int)
}

func (q *Queue) Empty() bool {
    return q.Size() == 0
}
