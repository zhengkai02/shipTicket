package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/quarkcms/quark-go/v2/internal/api"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"gorm.io/gorm"
	"time"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-22 17:08:11
*@Name account_service
*@Desc // 账号服务-保持所有账号在线
**/

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (s *AccountService) Start(ctx context.Context) error {
	for {
		cond := map[string]string{
			"status": "1",
		}
		var ret []*model.Account
		err := s.db.
			Model(&model.Account{}).
			Find(&ret, cond).
			Error
		if err != nil {
			log.Errorf("数据查询失败，err=[%v]", err)
		}
		// 检测账号是否掉线
		for _, account := range ret {
			go s.keepSession(account)
		}
		time.Sleep(5 * time.Second)
	}
}

func (s *AccountService) Stop(ctx context.Context) error {
	return nil
}

func (s *AccountService) keepSession(account *model.Account) {
	if err := api.CheckToken(account.Token); err != nil {
		resp, err := api.Login(account.Account, account.Password)
		if err != nil {
			log.Errorf("登录失败，err=[%v]", err)
			return
		}
		values := map[string]interface{}{
			"token":           resp.Data.Token,
			"account_type_id": resp.Data.AccountTypeId,
			"user_id":         resp.Data.UserId,
		}
		if err := s.db.
			Debug().
			Model(&model.Account{}).
			Where("account = ?", account.Account).
			Updates(values).Error; err != nil {
			log.Errorf("数据更新失败,err=[%v]", err)
		}
	}
}
