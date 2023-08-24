package resource

import (
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Task struct {
	resource.Template
}

// 初始化
func (p *Task) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "任务列表"
	// 模型
	p.Model = &model.Task{}
	// 分页
	p.PerPage = 10
	p.IndexOrder = "id asc"
	return p
}

// 只查询文章类型
//func (p *Article) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Task) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	users, _ := (&model.User{}).Options()
	ports, _ := (&model.Port{}).Options()
	portNames, _ := (&model.Port{}).Options1()

	return []interface{}{
		field.ID("id", "ID"),
		field.Select("departure_port_code", "出发港").SetRequired().HideFromIndex(true).SetOptions(ports),
		field.Select("departure_port_name", "出发港").SetOptions(portNames),
		field.Select("arrival_port_code", "到达港").SetRequired().HideFromIndex(true).SetOptions(ports),
		field.Select("arrival_port_name", "到达港").SetOptions(portNames),
		field.Date("departure_date", "出发日期").SetRequired(),
		field.Time("earliest_time", "最早出发时间"),
		field.Time("lastest_time", "最晚出发时间"),
		field.Number("passenger_num", "乘客数").SetRequired(),
		field.Number("vehicle_num", "车辆数").SetRequired(),
		// 单选模式
		field.Select("user_id", "联系人").SetOptions(users),

		//field.Editor("content", "内容").OnlyOnForms(),
		field.Datetime("create_time", "创建时间").HideWhenCreating(true),
		field.Datetime("update_time", "更新时间").HideWhenCreating(true),

		field.Switch("status", "状态").
			SetTrueValue("启动").
			SetFalseValue("停止").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Task) Searches(ctx *builder.Context) []interface{} {
	//options, _ := (&model.Line{}).TreeSelect(false)
	options, _ := (&model.User{}).Options()
	return []interface{}{
		searches.Select("user_id", "用户ID", options),
		searches.Status(),
		searches.DatetimeRange("create_time", "创建时间"),
	}
}

// 行为
func (p *Task) Actions(ctx *builder.Context) []interface{} {
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
