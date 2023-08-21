package resource

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/admin"
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"reflect"

	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/url"
)

type Order struct {
	resource.Template
}

// 初始化
func (p *Order) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "订单"
	// 模型
	p.Model = &model.Order{}
	// 分页
	p.PerPage = 10
	p.GET(resource.IndexPath, p.IndexRender)
	return p
}

// 只查询文章类型
//func (p *Line) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Order) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.Text("obtainTicketNum", "订单ID"),
		field.Text("passName", "乘客姓名"),
		field.Text("credentialNum", "证件号码"),
		field.Text("lineName", "航线"),
		field.Text("shipName", "型号"),
		field.Text("sailDate", "出发日期"),
		field.Text("busStartTime", "开车时间"),
		field.Text("sailTime", "开船时间"),
		field.Text("hxlxm", "类型"),
		field.Text("seatClassName", "仓位"),
		field.Text("ticketFee", "价格"),
		field.Switch("status", "状态").
			SetTrueValue("已出票").
			SetFalseValue("出票中").
			SetEditable(false).
			SetDefault(1),
	}
}

// 搜索
func (p *Order) Searches(ctx *builder.Context) []interface{} {
	options, _ := (&model.Port{}).Options()
	return []interface{}{
		searches.Select("orderId", "订单ID", options),
		//searches.Select("endPortNo", "到达港", options),
		//searches.Date("startDate", "日期"),
		//searches.Select("end_port_code", "到达", options),
		//searches.Input("start_port_code", "出发港口"),
		//searches.Input("end_port_code", "到达港口"),
		//searches.Status(),
	}
}

// 行为
func (p *Order) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.DetailLink(),
		//actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 列表页渲染
func (p *Order) IndexRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	//data := (&requests.IndexRequest{}).QueryData(ctx)
	data := p.orderList(ctx)

	// 组件渲染
	body := template.IndexComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

func (p *Order) orderList(ctx *builder.Context) interface{} {
	var lists []map[string]interface{}

	template := ctx.Template.(types.Resourcer)

	modelInstance := template.GetModel()

	model := db.Client.Model(modelInstance)

	// 搜索项
	searches := template.Searches(ctx)

	// 过滤项，预留
	filters := template.Filters(ctx)

	query := template.BuildIndexQuery(ctx, model, searches, filters, p.columnFilters(ctx), p.orderings(ctx))

	// 获取分页
	perPage := template.GetPerPage()
	if perPage == nil {
		query.Find(&lists)

		// 返回解析列表
		return p.performsList(ctx, lists)
	}

	// 不分页，直接返回lists
	if reflect.TypeOf(perPage).String() != "int" {
		query.Find(&lists)

		// 返回解析列表
		return p.performsList(ctx, lists)
	}

	var total int64
	var data map[string]interface{}
	page := 1
	querys := ctx.AllQuerys()
	if querys["search"] != nil {
		err := json.Unmarshal([]byte(querys["search"].(string)), &data)
		if err == nil {
			if data["current"] != nil {
				page = int(data["current"].(float64))
			}
			if data["pageSize"] != nil {
				perPage = int(data["pageSize"].(float64))
			}
		}
	}

	// 获取总数量
	//query.Count(&total)

	// 获取列表
	//query.Limit(perPage.(int)).Offset((page - 1) * perPage.(int)).Find(&lists)

	// 解析列表
	result := p.performsList(ctx, lists)
	log.Infof("查询成功")
	return map[string]interface{}{
		"currentPage": page,
		"perPage":     perPage,
		"total":       total,
		"items":       result,
		//"items": []interface{}{map[string]interface{}{"id": "zzzzzz", "LineName": "yyyyy", "ShipName": "xxxxxxxxx"}},
	}
}

// Get the column filters for the request.
func (p *Order) columnFilters(ctx *builder.Context) map[string]interface{} {
	querys := ctx.AllQuerys()
	var data map[string]interface{}
	if querys["filter"] == nil {
		return data
	}
	err := json.Unmarshal([]byte(querys["filter"].(string)), &data)
	if err != nil {
		return data
	}

	return data
}

// Get the orderings for the request.
func (p *Order) orderings(ctx *builder.Context) map[string]interface{} {
	querys := ctx.AllQuerys()
	var data map[string]interface{}
	if querys["sorter"] == nil {
		return data
	}
	err := json.Unmarshal([]byte(querys["sorter"].(string)), &data)
	if err != nil {
		return data
	}

	return data
}

// 处理列表
func (p *Order) performsList(ctx *builder.Context, lists []map[string]interface{}) []interface{} {
	var orderReq map[string]string
	searchParams := ctx.Querys
	for k, v := range searchParams {
		if k == "search" {
			err := json.Unmarshal([]byte(v.(string)), &orderReq)
			if err != nil {
				log.Error(err)
			}
		}
	}
	req := url.NewRequest()
	header := url.NewHeaders()
	for k, v := range generateHeader() {
		header.Set(k, v)
	}
	header.Set("session", "ssky_user_c006b156a69f4e6bbd567f7b128b7c87")
	req.Headers = header
	params := url.NewParams()
	params.Set("userId", "659965")
	params.Set("orderId", "216917408645766329")
	req.Params = params
	r, err := requests.Get(admin.OrderDetailAPI, req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := r.SimpleJson()
	b, _ := data.MarshalJSON()
	var resp admin.OrderDetailResp
	json.Unmarshal(b, &resp)
	var ret []interface{}
	for _, v := range resp.Data.OrderItemList {
		v.CredentialNum = "**************"
		v.Status = switchStatus(v.ItemState)
		v.TicketFee = float64(194)
		ret = append(ret, v)
	}
	header.Set("session", "ssky_user_fff823e2b9774dd4915d4c88aba31c01")
	header.Set("authentication", "1691743858652251")
	params.Set("userId", "652251")
	params.Set("orderId", "216917406395739885")
	req.Headers = header
	req.Params = params
	r, err = requests.Get(admin.OrderDetailAPI, req)
	if err != nil {
		fmt.Println(err)
	}
	var resp2 admin.OrderDetailResp
	data, _ = r.SimpleJson()
	b, _ = data.MarshalJSON()
	json.Unmarshal(b, &resp2)
	for _, v := range resp2.Data.OrderItemList {
		v.CredentialNum = "**************"
		v.Status = switchStatus(v.ItemState)
		v.TicketFee = float64(194)
		ret = append(ret, v)
	}
	return ret
}

func switchStatus(itemState int) int {
	switch itemState {
	case 2:
		return 1
	default:
		return 0
	}
}
