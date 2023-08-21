package resource

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/admin"
	"github.com/quarkcms/quark-go/v2/internal/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Ship struct {
	resource.Template
}

// 初始化
func (p *Ship) Init(ctx *builder.Context) interface{} {
	// 标题
	p.Title = "港口"
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

func (p *Ship) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.ID("id", "ID"),
		field.Text("lineName", "航线"),
		field.Text("shipName", "型号"),
		field.Text("sailDate", "日期"),
		field.Text("sailTime", "时间"),
		field.Datetime("startPortName", "出发港"),
		field.Datetime("endPortName", "到达港"),
		field.Datetime("className", "仓位"),
		field.Datetime("embarkPortName", "码头"),
		field.Datetime("pubCurrentCount", "余票"),
		field.Datetime("totalPrice", "价格"),
	}
}

// 搜索
func (p *Ship) Searches(ctx *builder.Context) []interface{} {
	options, _ := (&model.Port{}).Options()
	return []interface{}{
		searches.Select("startPortNo", "出发港", options),
		searches.Select("endPortNo", "到达港", options),
		searches.Date("startDate", "日期"),
		//searches.Select("end_port_code", "到达", options),
		//searches.Input("start_port_code", "出发港口"),
		//searches.Input("end_port_code", "到达港口"),
		//searches.Status(),
	}
}

// 行为
func (p *Ship) Actions(ctx *builder.Context) []interface{} {
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

// 列表页渲染
func (p *Ship) IndexRender(ctx *builder.Context) error {
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

func (p *Ship) shipList(ctx *builder.Context) interface{} {
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
func (p *Ship) columnFilters(ctx *builder.Context) map[string]interface{} {
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
func (p *Ship) orderings(ctx *builder.Context) map[string]interface{} {
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
func (p *Ship) performsList(ctx *builder.Context, lists []map[string]interface{}) []interface{} {
	ticketReq := &admin.TicketReq{
		StartPortNo: "1046",
		EndPortNo:   "1017",
		StartDate:   "2023-08-15",
	}
	params := ctx.Querys
	for k, v := range params {
		if k == "search" {
			err := json.Unmarshal([]byte(v.(string)), &ticketReq)
			if err != nil {
				log.Error(err)
			}
		}
	}
	reqBytes, _ := json.Marshal(ticketReq)
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, admin.EnqURL, bytes.NewReader(reqBytes))
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
	resBytes, _ := ioutil.ReadAll(resp.Body)
	var ticketResp admin.TicketResp
	if err := json.Unmarshal(resBytes, &ticketResp); err != nil {
		return nil
	}
	var ret []interface{}
	for i, v := range ticketResp.Data {
		for _, cls := range v.SeatClasses {
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
			ret = append(ret, item)
		}
	}
	return ret
}

func generateHeader() map[string]string {
	header := map[string]string{
		"Host":            `www.ssky123.com`,
		"Sec-Fetch-Site":  `same-origin`,
		"Accept-Language": `zh-CN,zh-Hans;q=0.9`,
		"Accept-Encoding": `gunzip, deflate, br`,
		"Sec-Fetch-Mode":  `cors`,
		"Origin":          `https://www.ssky123.com`,
		//"authentication":  `1684483703659965`,
		"authentication": `1691741017659965`,
		"User-Agent":     `Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.31(0x18001f37) NetType/WIFI Language/zh_CN`,
		"Referer":        `https://www.ssky123.com/online_booking/`,
		"Content-Length": `0`,
		"Connection":     `keep-alive`,
		"Sec-Fetch-Dest": `empty`,
		"Accept":         `application/json, text/plain, */*`,
		"Content-Type":   "application/json",
	}
	return header
}
