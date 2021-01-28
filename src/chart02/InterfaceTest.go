package chart02

// 接口
type Tree interface {
	Call() int8
}

// 必须有一个结构体
type TreeImpl struct {
	PhoneNumber int8
}

// 接口实现，go 当中是隐式实现
func (t TreeImpl) Call() int8 {
	return t.PhoneNumber * 10
}
