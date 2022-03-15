package gsa

type T struct {
	Val         int
	Left, Right *T
}

type FrameType int // Stack frame types
const (
	Process FrameType = iota
	Emit
)

// Stack frames is one of the two types plus
// the data
type StackFrame struct {
	action FrameType
	v      *T
}

func process(v *T) StackFrame {
	return StackFrame{Process, v}
}

func emit(v *T) StackFrame {
	return StackFrame{Emit, v}
}

type Stack []StackFrame

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(frame StackFrame) {
	*s = append(*s, frame)
}

func (s *Stack) pop() StackFrame {
	if s.isEmpty() {
		panic("Don't pop an empty stack")
	}
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v
}

// Do an in-order traversal of v and output the
// values in the tree.
func InOrder(v *T) []int {
	if v == nil {
		return []int{}
	}

	stack := Stack{process(v)}
	res := []int{}
	for !stack.isEmpty() {
		frame := stack.pop()
		v = frame.v
		switch frame.action {
		case Emit:
			res = append(res, v.Val)
		case Process:
			if v.Right != nil {
				stack.push(process(v.Right))
			}
			stack.push(emit(v))
			if v.Left != nil {
				stack.push(process(v.Left))
			}
		}
	}

	return res
}

type Queue []*T

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) enqueue(v *T) {
	if v != nil {
		*q = append(*q, v)
	}
}

func (q *Queue) dequeue() *T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

// Do a breadth-order traversal of v and output the
// values in the tree.
func BfOrder(v *T) []int {
	if v == nil {
		return []int{}
	}

	q := Queue{v}
	res := []int{}
	for !q.isEmpty() {
		v := q.dequeue()
		res = append(res, v.Val)
		q.enqueue(v.Left)
		q.enqueue(v.Right)
	}

	return res
}
