package main

import (
	"fmt"
)

type MyQueue struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	p := &MyQueue{}
	return *p
}

func (this *MyQueue) Push_to_top(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyQueue) Pop_from_top() int {
	x := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return x
}

func (this *MyQueue) Size() int {
	return len(this.Queue)
}

func (this *MyQueue) Is_empty() bool {
	return len(this.Queue) == 0
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Push_to_top(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp := &MyQueue{}
	size := this.Size()
	for i := 0; i < size-1; i++ {
		x := this.Pop_from_top()
		tmp.Push_to_top(x)
	}
	ans := this.Pop_from_top()
	tmp_size := tmp.Size()
	for i := 0; i < tmp_size; i++ {
		x := tmp.Pop_from_top()
		this.Push_to_top(x)
	}
	return ans
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	tmp := &MyQueue{}
	size := this.Size()
	for i := 0; i < size-1; i++ {
		x := this.Pop_from_top()
		tmp.Push_to_top(x)
	}
	ans := this.Pop_from_top()
	this.Push_to_top(ans)
	tmp_size := tmp.Size()
	for i := 0; i < tmp_size; i++ {
		x := tmp.Pop_from_top()
		this.Push_to_top(x)
	}
	return ans
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Is_empty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	fmt.Println(queue.Size())
	queue.Push(2)
	fmt.Println(queue.Size())
	queue.Push(3)
	fmt.Printf("%#v\n", queue)
	fmt.Println("######################")
	fmt.Println(queue.Peek())
}
