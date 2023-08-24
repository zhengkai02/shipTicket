package service

import (
	"context"
	"github.com/DamnWidget/goqueue"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/data"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"gorm.io/gorm"
	"time"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-21 11:27:40
*@Name agent_service
*@Desc // 代理人服务
**/

type AgentService struct {
	db    *gorm.DB
	queue *goqueue.Queue
}

func NewAgentService(db *gorm.DB) *AgentService {
	return &AgentService{
		db:    db,
		queue: data.NewGlobalQueue(),
	}
}

func (a AgentService) Start(ctx context.Context) error {
	for {
		cond := map[string]string{
			"status": "1",
		}
		var ret []*model.Task
		err := a.db.
			Model(&model.Task{}).
			Find(&ret, cond).
			Error
		if err != nil {
			log.Errorf("数据查询失败，err=[%v]", err)
		}
		for _, task := range ret {
			log.Debugf("任务[%v]放入队列", task.Id)
			if err := a.queue.Push(task); err != nil {
				log.Errorf("任务加入队列失败,err=[%v]", err)
				continue
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func (a AgentService) Stop(ctx context.Context) error {
	return nil
}
