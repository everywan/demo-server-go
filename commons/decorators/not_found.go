package decorators

// 典型场景: 当数据库里找不到记录时, 不返回 error,
// 通过 Resp 添加 EmptyDecorator 实现判断为空的能力.
type EmptyDecorator interface {
	IsExist() bool
}

type empty struct {
	isExist bool
}

func (e *empty) IsExist() bool {
	return e.isExist
}

var _ EmptyDecorator = new(empty)

func NewEmpty(exist bool) *empty {
	return &empty{
		isExist: exist,
	}
}
