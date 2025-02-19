type CustomStack struct {
	Capacity int
	Stack    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{Stack: make([]int, 0, maxSize), Capacity: maxSize}
}

func (this *CustomStack) Push(x int) {
	if len(this.Stack) >= this.Capacity {
		return
	}
	this.Stack = append(this.Stack, x)
}

func (this *CustomStack) Pop() int {
	if len(this.Stack) <= 0 {
		return -1
	}
	num := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return num
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < len(this.Stack) && i < k; i++ {
		this.Stack[i] += val
	}
}
