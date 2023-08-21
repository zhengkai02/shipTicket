package resource

import (
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Line struct {
	resource.Template
}

// 初始化
func (p *Line) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "港口"
	// 模型
	p.Model = &model.Line{}
	// 分页
	p.PerPage = 10
	p.IndexOrder = "id asc"
	p.GET(resource.IndexPath, p.IndexRender)
	return p
}

// 只查询文章类型
//func (p *Line) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Line) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("start_port_code", "出发地编码").
			SetRules([]*rule.Rule{
				rule.Required(true, "出发地编码必须填写"),
			}).SetRequired(),

		field.Text("end_port_code", "目的地编码").SetRequired(),
		field.Text("start_port_name", "出发地"),
		field.Text("end_port_name", "目的地"),
		//field.Editor("content", "内容").OnlyOnForms(),
		field.Datetime("create_time", "创建时间"),
		field.Datetime("update_time", "更新时间"),
		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Line) Searches(ctx *builder.Context) []interface{} {
	options, _ := (&model.Port{}).Options()
	return []interface{}{
		searches.Select("start_port_code", "出发", options),
		searches.Select("end_port_code", "到达", options),
		searches.Input("start_port_code", "出发港口"),
		searches.Input("end_port_code", "到达港口"),
		searches.Status(),
	}
}

// 行为
func (p *Line) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.EditLink(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
