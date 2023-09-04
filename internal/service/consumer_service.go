package service

import (
	"context"
	"fmt"
	"github.com/DamnWidget/goqueue"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/api"
	"github.com/quarkcms/quark-go/v2/internal/data"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db    *gorm.DB
}

func NewConsumerService(db *gorm.DB) *ConsumerService {
	return &ConsumerService{
		queue: data.NewGlobalQueue(),
		db:    db,
	}
}

func (s *ConsumerService) Start(ctx context.Context) error {
	for {
		item := s.queue.Pop()
		if item == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Debugf("消费者收到任务消息: %v", item)
		task, ok := item.(*model.Task)
		if !ok {
			log.Errorf("消息格式错误：mgs=[%v]", task)
			continue
		}
		go func(s *ConsumerService, t *model.Task) {
			task.CreateTime = time.Now()
			if err := s.process(task); err != nil {
				log.Errorf("任务处理失败，err=[%v]", err.Error())
				return
			}
			log.Debugf("任务处理成功,耗时=[%v]", time.Since(task.CreateTime))
		}(s, task)

	}
}

func (s *ConsumerService) Stop(ctx context.Context) error {
	return nil
}

func (s *ConsumerService) process(t *model.Task) error {
	log.Infof("处理任务：[%v-%v-%v %v-%v]", t.DeparturePortName, t.ArrivalPortName, t.DepartureDate.Format(time.DateOnly), t.EarliestTime, t.LastestTime)
	switch true {
	case t.VehicleNum > 0:
		// 摆渡车
		ferryTicketList, err := api.FerryTicketList(t.DeparturePortCode, t.ArrivalPortCode, t.DepartureDate.Format(time.DateOnly))
		if err != nil {
			log.Errorf("查询客车票失败，err=[%v]", err)
		}
		for _, ticket := range ferryTicketList {
			// 时间筛选
			sailTime, _ := time.Parse("15:04", ticket.SailTime)
			earliestTime, _ := time.Parse("15:04:05", t.EarliestTime)
			latestTime, _ := time.Parse("15:04:05", t.LastestTime)
			if sailTime.Before(earliestTime) || sailTime.After(latestTime) {
				continue
			}
			for _, dsc := range ticket.DriverSeatClass {
				if dsc.PubCurrentCount >= t.VehicleNum {
					if t.PassengerNum > 0 {
						for _, cls := range ticket.SeatClasses {
							if cls.PubCurrentCount >= t.PassengerNum {
								msg := fmt.Sprintf("检测到摆渡车余票【%s-%s-%s %s %v ￥%v】【%v】张", ticket.StartPortName, ticket.EndPortName, ticket.SailDate, ticket.SailTime, cls.ClassName, cls.TotalPrice, cls.PubCurrentCount)
								api.SendMsg("trip", msg)
								log.Warnf(msg)
								return nil
							}
						}
					} else {
						msg := fmt.Sprintf("检测到摆渡车余票【%s-%s-%s %s %v ￥%v】【%v】张", ticket.StartPortName, ticket.EndPortName, ticket.SailDate, ticket.SailTime, dsc.ClassName, dsc.TotalPrice, dsc.PubCurrentCount)
						api.SendMsg("trip", msg)
						log.Warnf(msg)
					}
				}

			}
		}
	case t.VehicleNum == 0:
		// 查询航班
		ticketList, err := api.ShipTicketList(t.DeparturePortCode, t.ArrivalPortCode, t.DepartureDate.Format(time.DateOnly))
		if err != nil {
			log.Errorf("航班查询失败，err=[%v]", err)
			return err
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
					orderId := fmt.Sprintf("B%v%v%v%v", t.DeparturePortCode, t.ArrivalPortCode, t.DepartureDate.Format("20060102"), t.UserId)

					var count int64
					cond := map[string]string{
						"order_id": orderId,
					}
					err := s.db.Model(&model.Order{}).Where(cond).Count(&count).Error
					if count > 0 {
						break
					}
					departureTimeStr := ticket.SailDate + ticket.SailTime
					departureTime, _ := time.ParseInLocation("2006/01/0215:04", departureTimeStr, time.Local)
					order := &model.Order{
						ID:                0,
						OrderID:           orderId,
						SuplierOrderID:    "",
						DeparturePortName: ticket.StartPortName,
						ArrivalPortName:   ticket.EndPortName,
						DepartureTime:     departureTime,
						PassengerNum:      t.PassengerNum,
						VehicleNum:        t.VehicleNum,
						CreateTime:        time.Now(),
						UpdateTime:        time.Now(),
					}
					// 不处理冲突
					err = s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(order).Error
					if err != nil {
						return err
					}
					//s.db.Clauses(clause.OnConflict{
					//	Columns:   []clause.Column{{Name: "id"}},
					//	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
					//}).Create(order)
					msg := fmt.Sprintf("检测到旅客余票【%s-%s-%s %s %v ￥%v】【%v】张", ticket.StartPortName, ticket.EndPortName, ticket.SailDate, ticket.SailTime, cls.ClassName, cls.TotalPrice, cls.PubCurrentCount)
					api.SendMsg("trip", msg)
					log.Info(msg)
				}
			}
		}
	}
	return nil
}
