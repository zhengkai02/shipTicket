package service

import "github.com/quarkcms/quark-go/v2/internal/admin/resource"

// 注册服务
var Provider = []interface{}{
	&resource.Port{},
	&resource.Line{},
	&resource.Ship{},
	&resource.Ferry{},
	&resource.Order{},
	&resource.Account{},
	&resource.Task{},
	&resource.User{},
	&resource.Passenger{},
}
