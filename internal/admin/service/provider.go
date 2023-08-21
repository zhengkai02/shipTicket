package service

import "github.com/quarkcms/quark-go/v2/internal/admin/resource"

// 注册服务
var Provider = []interface{}{
	&resource.Article{},
	&resource.Line{},
	&resource.Ship{},
	&resource.Order{},
}
