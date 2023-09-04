package resource

import (
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type User struct {
	resource.Template
}

// 初始化
func (p *User) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "联系人"
	// 模型
	p.Model = &model.User{}
	// 分页
	p.PerPage = 10
	p.IndexOrder = "id asc"
	return p
}

// 只查询文章类型
//func (p *Article) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *User) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.ID("id", "ID"),
		field.Number("user_id", "用户标识").SetRequired(),
		field.Text("name", "用户名").SetRequired(),
		field.Text("phone", "电话").SetRequired(),

		//field.Editor("content", "内容").OnlyOnForms(),
		field.Datetime("create_time", "创建时间").HideWhenCreating(true),
		field.Datetime("update_time", "更新时间"),
	}
}

// 搜索
func (p *User) Searches(ctx *builder.Context) []interface{} {
	//options, _ := (&model.Line{}).TreeSelect(false)

	return []interface{}{
		searches.Input("name", "用户名"),
		searches.Status(),
		searches.DatetimeRange("create_time", "创建时间"),
	}
}

// 行为
func (p *User) Actions(ctx *builder.Context) []interface{} {
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
