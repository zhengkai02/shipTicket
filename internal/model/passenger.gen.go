// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePassenger = "t_passenger"

// Passenger mapped from table <t_passenger>
type Passenger struct {
	Id         int64     `gorm:"column:id;type:int(8);primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	IdCard     string    `gorm:"column:id_card;type:varchar(255);not null" json:"id_card"`
	CarNo      string    `gorm:"column:car_no;type:varchar(255);not null" json:"car_no"`
	UserId     int64     `gorm:"column:user_id;type:int(8)" json:"user_id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:2023-07-24 00:00:00" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:2023-07-24 00:00:00" json:"update_time"`
}

// TableName Passenger's table name
func (*Passenger) TableName() string {
	return TableNamePassenger
}