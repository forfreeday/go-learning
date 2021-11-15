package P02interface

// 接口
type Tree interface {
	Call() int8
}

// 必须有一个结构体
type TreeImpl struct {
	PhoneNumber int8
}

// 必须有一个结构体
type TreeImpl2 struct {
	PhoneNumber int8
}

// 接口实现，go 当中是隐式实现
func (t TreeImpl) Call() int8 {
	return t.PhoneNumber * 10
}

func (t TreeImpl2) Call() int8 {
	return t.PhoneNumber * 10
}
