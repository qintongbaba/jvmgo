package rtda

//栈帧
type Frame struct {
	lower        *Frame
	localvars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localvars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localvars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
