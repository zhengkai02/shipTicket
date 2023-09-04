package resource

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/admin"
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"io"
	"net/http"
	"reflect"
	"time"
)

type Ferry struct {
	resource.Template
}

// 初始化
func (p *Ferry) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "小客车及随行人员"
	p.SubTitle = "子标题"
	// 模型
	p.Model = &model.Ship{}
	// 分页
	p.PerPage = 10
	p.GET(resource.IndexPath, p.IndexRender)
	return p
}

// 只查询文章类型
//func (p *Line) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
//	return query.Debug().Where("status", "1")
//}

func (p *Ferry) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.ID("id", "ID"),
		//field.Text("lineName", "航线"),
		field.Text("shipName", "名称"),
		field.Text("clxm", "型号"),
		field.Text("sailDate", "日期"),
		field.Text("sailTime", "时间"),
		field.Text("startPortName", "出发港").SetDefault(1028),
		field.Text("endPortName", "到达港").SetDefault(1017),

		field.Text("className", "车辆舱位"),
		field.Number("pubCurrentCount", "余票"),
		field.Text("passengerClassName", "旅客舱位"),
		field.Number("pubCurrentCount", "旅客余票"),
		field.Text("totalPrice", "价格").SetEditable(true),
	}
}

// 搜索
func (p *Ferry) Searches(ctx *builder.Context) []interface{} {
	options, _ := (&model.Port{}).Options()
	res := []interface{}{
		searches.Select("startPortNo", "出发港", options),
		searches.Select("endPortNo", "到达港", options),
		searches.Date("startDate", "日期"),
		searches.Select("class_name", "仓位", []*selectfield.Option{
			{Label: "上舱", Value: "上舱"},
			{Label: "中舱", Value: "中舱"},
			{Label: "下舱", Value: "下舱"},
		}),
		searches.Select("interval", "时间段", []*selectfield.Option{
			{Label: "上午", Value: "am"},
			{Label: "下午", Value: "pm"},
		}),
		searches.Select("clxm", "型号", []*selectfield.Option{
			{Label: "客滚船", Value: "客滚船"},
			{Label: "高速客船", Value: "高速客船"},
			{Label: "游轮客船", Value: "游轮客船"}, //	常规客船
			{Label: "常规客船", Value: "常规客船"},
		}),
		//searches.Input("start_port_code", "出发港口"),
		//searches.Input("end_port_code", "到达港口"),
		//searches.Status(),
	}
	return res
}

// 行为
func (p *Ferry) Actions(ctx *builder.Context) []interface{} {
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
		actions.DetailLink(),
	}
}

// 列表页渲染
func (p *Ferry) IndexRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	//data := (&requests.IndexRequest{}).QueryData(ctx)
	data := p.shipList(ctx)

	// 组件渲染
	body := template.IndexComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

func (p *Ferry) shipList(ctx *builder.Context) interface{} {
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
	// 解析列表
	result := p.performsList(ctx, lists)
	log.Infof("查询成功")
	return map[string]interface{}{
		"currentPage": page,
		"perPage":     perPage,
		"total":       total,
		"items":       result,
		//"items": []api{}{map[string]api{}{"id": "zzzzzz", "LineName": "yyyyy", "ShipName": "xxxxxxxxx"}},
	}
}

// Get the column filters for the request.
func (p *Ferry) columnFilters(ctx *builder.Context) map[string]interface{} {
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
func (p *Ferry) orderings(ctx *builder.Context) map[string]interface{} {
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
func (p *Ferry) performsList(ctx *builder.Context, lists []map[string]interface{}) []interface{} {
	ticketReq := &admin.TicketReq{
		StartPortNo: "1028",
		EndPortNo:   "1010",
		StartDate:   time.Now().AddDate(0, 0, 1).Format(time.DateOnly),
	}
	ticketFilter := &admin.TicketFilter{}
	if params, ok := ctx.Querys["search"]; ok {
		err := json.Unmarshal([]byte(params.(string)), &ticketReq)
		if err != nil {
			log.Error(err)
		}
		date, _ := time.ParseInLocation(time.DateTime, ticketReq.StartDate, time.Local)
		if date.After(time.Now()) {
			ticketReq.StartDate = date.Format(time.DateOnly)
		}
		err = json.Unmarshal([]byte(params.(string)), &ticketFilter)
		if err != nil {
			log.Error(err)
		}
	}
	reqBytes, _ := json.Marshal(ticketReq)
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, admin.FerryEnqURL, bytes.NewReader(reqBytes))
	if err != nil {
		return nil
	}
	for k, v := range generateHeader() {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode > 201 {
		return nil
	}
	defer resp.Body.Close()
	resBytes, _ := io.ReadAll(resp.Body)
	var ticketResp admin.TicketResp
	if err := json.Unmarshal(resBytes, &ticketResp); err != nil {
		return nil
	}
	var ret []interface{}
	for i, v := range ticketResp.Data {
		switchInterval := func(sailTime string) string {
			startTime, _ := time.Parse("15:04", sailTime)
			standardTime, _ := time.Parse("15:04", "12:00")
			if startTime.Before(standardTime) {
				return "am"
			}
			return "pm"
		}
		if len(ticketFilter.Interval) > 0 && switchInterval(v.SailTime) != ticketFilter.Interval {
			continue
		}
		if len(ticketFilter.Clxm) > 0 && v.Clxm != ticketFilter.Clxm {
			continue
		}
		for _, cls := range v.DriverSeatClass {
			if len(ticketFilter.ClassName) > 0 && cls.ClassName != ticketFilter.ClassName {
				continue
			}
			v.ID = i
			var item map[string]interface{}
			dataBytes, _ := json.Marshal(v)
			json.Unmarshal(dataBytes, &item)
			clsMap := map[string]interface{}{}
			c, _ := json.Marshal(cls)
			json.Unmarshal(c, &clsMap)
			for k, v := range clsMap {
				item[k] = v
			}
			if len(v.SeatClasses) > 0 {
				for _, cls := range v.SeatClasses {
					tmpItem := make(map[string]interface{})
					for k, v := range item {
						tmpItem[k] = v
					}
					tmpItem["passengerClassName"] = cls.ClassName
					tmpItem["reseat"] = cls.PubCurrentCount
					ret = append(ret, tmpItem)
				}
				continue
			}
			ret = append(ret, item)
		}
	}
	return ret
}
