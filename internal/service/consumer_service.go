package service

import (
	"context"
	"fmt"
	"github.com/DamnWidget/goqueue"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/api"
	"github.com/quarkcms/quark-go/v2/internal/data"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"time"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-22 9:40:44
*@Name consumer_service
*@Desc // 消费者服务
**/

type ConsumerService struct {
	queue *goqueue.Queue
}

func NewConsumerService() *ConsumerService {
	return &ConsumerService{
		queue: data.NewGlobalQueue(),
	}
}

func (s *ConsumerService) Start(ctx context.Context) error {
	for {
		item := s.queue.Pop()
		if item == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Infof("消费者收到任务消息: %v", item)
		task, ok := item.(*model.Task)
		if !ok {
			log.Errorf("消息格式错误：mgs=[%v]", task)
			continue
		}
		go func(t *model.Task) {
			task.CreateTime = time.Now()
			if err := process(task); err != nil {
				log.Errorf("任务处理失败，err=[%v]", err.Error())
				return
			}
			log.Infof("任务处理成功,耗时=[%v]", time.Since(task.CreateTime))
		}(task)

	}
}

func (s *ConsumerService) Stop(ctx context.Context) error {
	return nil
}

func process(t *model.Task) error {
	log.Infof("处理任务：[%v-%v-%v]", t.DepaturePortName, t.ArrvalPortName, t.EarliestTime)
	// 查询航班
	ticketList, err := api.ShipTicketList(t.DepaturePortCode, t.ArrivalPortCode, t.DepartureDate.Format(time.DateOnly))
	if err != nil {
		log.Errorf("航班查询失败，err=[%v]", err)
		return err
	}
	// 摆渡车
	if t.VehicleNum > 0 {

	}
	// 根据时间区间过滤航班
	for _, ticket := range ticketList {
		// 时间筛选
		sailTime, _ := time.Parse("15:04", ticket.SailTime)
		earliestTime, _ := time.Parse("15:04:05", t.EarliestTime)
		latestTime, _ := time.Parse("15:04:05", t.LastestTime)
		if sailTime.Before(earliestTime) || sailTime.After(latestTime) {
			continue
		}
		for _, cls := range ticket.SeatClasses {
			if cls.PubCurrentCount >= t.PassengerNum {
				msg := fmt.Sprintf("检测到余票【%s-%s-%s %s %v ￥%v】【%v】张", ticket.StartPortName, ticket.EndPortName, ticket.SailDate, ticket.SailTime, cls.ClassName, cls.TotalPrice, cls.PubCurrentCount)
				api.SendMsg("trip", msg)
				log.Info(msg)
			}
		}
	}

	return nil
}
