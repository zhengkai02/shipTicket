package server

import (
	"context"
	"github.com/quarkcms/quark-go/v2/internal/service"
	"gorm.io/gorm"
)

/**
*@Auther kaikai.zheng
*@Date 2022-12-29 11:27:28
*@Name agent
*@Desc 代理人 agent
**/

// WatchServer 后台监听服务
type WatchServer struct {
	services []AgentService
}

// Agent 代理人服务
type AgentService interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// NewWatchServer 新建后台监听服务
func NewWatchServer(db *gorm.DB) *WatchServer {
	ws := &WatchServer{}
	// kafka消费者服务
	agentService := service.NewAgentService(db)
	ws.RegisterService(agentService)

	//rs := service.NewRedisService(pricingRule)
	//cs.RegisterService(rs)

	cs := service.NewConsumerService()
	ws.RegisterService(cs)
	return ws
}

func (s *WatchServer) RegisterService(as AgentService) {
	s.services = append(s.services, as)
}

// Start 启动后台监听服务服务
func (c *WatchServer) Start(ctx context.Context) error {
	for _, s := range c.services {
		go s.Start(ctx)
	}
	return nil
}

// Stop 停止后台监听服务服务
func (c *WatchServer) Stop(ctx context.Context) error {
	for _, s := range c.services {
		s.Stop(ctx)
	}
	return nil
}
