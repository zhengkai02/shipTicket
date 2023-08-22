package data

import (
	"github.com/DamnWidget/goqueue"
	"sync"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-22 9:31:49
*@Name queue
*@Desc // 全局队列
**/

var (
	GlobalQueue *goqueue.Queue
	once        sync.Once
)

func NewGlobalQueue() *goqueue.Queue {
	once.Do(func() {
		if GlobalQueue == nil {
			q := goqueue.New()
			GlobalQueue = q
		}
	})
	return GlobalQueue
}
