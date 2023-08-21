package service

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
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
	db *gorm.DB
}

func NewAgentService(db *gorm.DB) *AgentService {
	return &AgentService{
		db: db,
	}
}

func (a AgentService) Start(ctx context.Context) error {
	for {
		cond := map[string]string{
			"status": "1",
		}
		var ret []*model.Task
		err := a.db.
			Debug().
			Model(&model.Task{}).
			Find(&ret, cond).
			Error
		if err != nil {
			panic(err)
		}
		for _, task := range ret {
			fmt.Println(task.UserID, task.Passengers)
			log.Infof("任务[%v]放入队列", task.ID)
		}
		time.Sleep(5 * time.Second)
	}

	return nil
}

func (a AgentService) Stop(ctx context.Context) error {
	return nil
}
