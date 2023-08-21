package search

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-09 13:48:36
*@Name datetime
*@Desc //
**/

type DatetimeField struct {
	searches.DatetimeRange
}

// 日期时间
func Datetime(column string, name string) *DatetimeField {
	field := &DatetimeField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *DatetimeField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}
