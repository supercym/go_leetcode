type MinStack struct {
    Stack []int
    Min_val int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    p := &MinStack{}
    return *p
}


func (this *MinStack) Push(x int)  {  
    this.Stack = append(this.Stack, x)
    if len(this.Stack) == 1 {
        this.Min_val = x
    }
    if this.Min_val > x {
        this.Min_val = x
    }
}


func (this *MinStack) Pop()  {
    if this.Stack[len(this.Stack)-1] == this.Min_val {
        this.Min_val = this.Stack[0]
        for i := 0; i < len(this.Stack)-1; i++ {
            if this.Min_val > this.Stack[i] {
                this.Min_val = this.Stack[i]
            }
        }
    }
    this.Stack = this.Stack[:len(this.Stack)-1]
    
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.Min_val
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
 