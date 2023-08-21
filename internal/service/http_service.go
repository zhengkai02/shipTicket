package service

import (
	"context"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	myservice "github.com/quarkcms/quark-go/v2/internal/admin/service"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/install"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/middleware"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service"
	toolservice "github.com/quarkcms/quark-go/v2/pkg/app/tool/service"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-21 16:39:35
*@Name http
*@Desc //
**/

type HttpService struct {
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (a HttpService) Start(ctx context.Context) error {
	// 服务
	var providers []interface{}
	// 加载后台服务
	providers = append(providers, service.Providers...)

	// 加载自定义后台服务
	providers = append(providers, myservice.Provider...)

	// 加载工具服务
	providers = append(providers, toolservice.Providers...)
	// 数据库配置信息
	dsn := "root:Perr78!!@tcp(localhost:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{

		// JWT加密密串
		AppKey: "123456",

		// 加载服务
		Providers: providers,

		// 数据库配置
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// WEB根目录
	b.Static("/", "./web/app")

	// 自动构建数据库、拉取静态文件
	install.Handle()

	// 后台中间件
	b.Use(middleware.Handle)

	// 响应Get请求
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 开启Debug模式
	b.Echo().Debug = true

	// 日志文件位置
	//f, _ := os.OpenFile("./app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	// 记录日志
	b.Echo().Logger.SetOutput(io.MultiWriter(os.Stdout))

	// 日志中间件
	//b.Echo().Use(echomiddleware.Logger())

	// 崩溃后自动恢复
	b.Echo().Use(echomiddleware.Recover())

	// 启动服务
	b.Run(":3000")
	return nil
}

func (a HttpService) Stop(ctx context.Context) error {
	return nil
}
