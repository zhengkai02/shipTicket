package resource

import (
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Passenger struct {
	resource.Template
}

// 初始化
func (p *Passenger) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "乘客信息"
	// 模型
	p.Model = &model.Passenger{}
	// 分页
	p.PerPage = 10
	p.IndexOrder = "id asc"
	return p
}

// 只查询文章类型
//func (p *Article) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Passenger) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	users, _ := (&model.User{}).Options()
	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "用户名").SetRequired(),
		field.Text("id_card", "身份证").SetRequired(),
		field.Text("car_no", "车牌号"),
		field.Select("user_id", "联系人").SetOptions(users),

		//field.Editor("content", "内容").OnlyOnForms(),
		field.Datetime("create_time", "创建时间").HideWhenCreating(true),
		field.Datetime("update_time", "更新时间").HideWhenCreating(true),
	}
}

// 搜索
func (p *Passenger) Searches(ctx *builder.Context) []interface{} {
	//options, _ := (&model.Line{}).TreeSelect(false)
	users, _ := (&model.User{}).Options()
	account, _ := (&model.Account{}).Options()
	users = append(users, account...)
	return []interface{}{
		searches.Input("name", "用户名"),
		searches.Select("user_id", "联系人", users),
		searches.Status(),
		searches.DatetimeRange("create_time", "创建时间"),
	}
}

// 行为
func (p *Passenger) Actions(ctx *builder.Context) []interface{} {
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
