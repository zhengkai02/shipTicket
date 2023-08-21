package types

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Drawer interface {
	Actioner

	// 宽度
	GetWidth() int

	// 关闭时销毁 Drawer 里的子元素
	GetDestroyOnClose() bool

	// 内容
	GetBody(ctx *builder.Context) interface{}

	// 弹窗行为
	GetActions(ctx *builder.Context) []interface{}
}
