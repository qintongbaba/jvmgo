package rtda

//局部变量空间
type Slot struct {
	num int32   //存放整数
	ref *Object //存放引用
}
