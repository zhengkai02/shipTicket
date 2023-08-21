package action

import (
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type CreateLinkAction struct {
	actions.Link
}

// 创建-跳转类型
func CreateLink() *CreateLinkAction {
	return &CreateLinkAction{}
}

// 初始化
func (p *CreateLinkAction) Init(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 文字
	p.Name = "创建" + template.GetTitle()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 跳转链接
func (p *CreateLinkAction) GetHref(ctx *builder.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/index", "/create", -1)
}

/**
SetOnlyOnIndex 只在列表页展示
SetExceptOnIndex 除了列表页外展示
SetOnlyOnForm 只在表单页展示
SetExceptOnForm 除了表单页外展示
SetOnlyOnFormExtra 只在表单页右上角自定义区域展示
SetExceptOnFormExtra 除了表单页右上角自定义区域外展示
SetOnlyOnDetail 只在详情页展示
SetExceptOnDetail 除了详情页外展示
SetOnlyOnDetailExtra 只在详情页右上角自定义区域展示
SetExceptOnDetailExtra 除了详情页右上角自定义区域外展示
SetOnlyOnIndexTableRow 在表格行内展示
SetExceptOnIndexTableRow 除了表格行内外展示
SetOnlyOnIndexTableAlert 在表格多选弹出层展示
SetExceptOnIndexTableAlert 除了表格多选弹出层外展示
SetShowOnIndex 在列表页展示
SetShowOnForm 在表单页展示
SetShowOnFormExtra 在表单页右上角自定义区域展示
SetShowOnDetail 在详情页展示
SetShowOnDetailExtra 在详情页右上角自定义区域展示
SetShowOnIndexTableRow 在表格行内展示
SetShowOnIndexTableAlert 在多选弹出层展示


*/
