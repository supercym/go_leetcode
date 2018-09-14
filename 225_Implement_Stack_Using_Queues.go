type MyStack struct {
    Stack []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    p := &MyStack{}
    return *p
}

func (this *MyStack) Push_to_back(x int) {
    this.Stack = append(this.Stack, x)
}

func (this *MyStack) Pop_from_front()int {
    head := this.Stack[0]
    this.Stack = this.Stack[1:len(this.Stack)]
    return head
}

func (this *MyStack) Size()int {
    return len(this.Stack)
}

func (this *MyStack) Is_empty()bool {
    return len(this.Stack) == 0
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Push_to_back(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    size := this.Size()
    for i := 0; i < size-1; i++ {
        x := this.Pop_from_front()
        this.Push_to_back(x)
    }
    return this.Pop_from_front()
}


/** Get the top element. */
func (this *MyStack) Top() int {
    x := this.Pop()
    this.Push_to_back(x)
    return x
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.Is_empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
 