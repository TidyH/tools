package main

type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	var NewStack MinStack

	return NewStack
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)

	if len(this.Stack) == 1 {
		this.Min = append(this.Min, val)
	} else if val <= this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, val)
	}
}

func (this *MinStack) Pop() {
	pop := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]

	if this.Min[len(this.Min)-1] == pop {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	println("155. Min Stack")

	// obj := Constructor()
	// obj.Push(val)
	// obj.Pop()
	// param_3 := obj.Top()
	// param_4 := obj.GetMin()

	obj := Constructor()
	obj.Push(-10)
	obj.Push(2)
	println(obj.GetMin())
	obj.Push(-20)
	println(obj.GetMin())
	obj.Pop()
	println(obj.GetMin())

}
