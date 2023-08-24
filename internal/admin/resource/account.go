package resource

import (
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Account struct {
	resource.Template
}

// 初始化
func (p *Account) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "账号管理"
	// 模型
	p.Model = &model.Account{}
	// 分页
	p.PerPage = 10
	p.IndexOrder = "id asc"
	return p
}

// 只查询文章类型
//func (p *Article) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Account) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("account", "账号").
			SetRules([]*rule.Rule{
				rule.Required(true, "账号名必填"),
			}).SetRequired(),

		field.Password("password", "密码").SetCopyable(true),
		field.Text("account_type_id", "账号类型").HideWhenCreating(true),
		field.Text("user_id", "用户标识").HideWhenCreating(true),
		field.Text("token", "token").SetCopyable(true).HideWhenCreating(true),

		//field.Editor("content", "内容").OnlyOnForms(),
		field.Datetime("create_time", "创建时间").HideWhenCreating(true),
		field.Datetime("update_time", "更新时间").HideWhenCreating(true),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Account) Searches(ctx *builder.Context) []interface{} {
	//options, _ := (&model.Line{}).TreeSelect(false)

	return []interface{}{
		searches.Input("account", "账号"),
		searches.Status(),
		searches.DatetimeRange("create_time", "创建时间"),
	}
}

// 行为
func (p *Account) Actions(ctx *builder.Context) []interface{} {
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
