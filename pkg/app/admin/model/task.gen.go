// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTask = "t_task"

// Task mapped from table <t_task>
type Task struct {
	ID               int64     `gorm:"column:id;type:bigint(64) unsigned;primaryKey;autoIncrement:true" json:"id"`
	DepaturePortCode string    `gorm:"column:depature_port_code;type:varchar(255);not null" json:"depature_port_code"`
	DepaturePortName string    `gorm:"column:depature_port_name;type:varchar(255);not null" json:"depature_port_name"`
	ArrivalPortCode  string    `gorm:"column:arrival_port_code;type:varchar(255);not null" json:"arrival_port_code"`
	ArrvalPortName   string    `gorm:"column:arrval_port_name;type:varchar(255);not null" json:"arrval_port_name"`
	EarliestTime     time.Time `gorm:"column:earliest_time;type:datetime;not null" json:"earliest_time"`
	LatestTime       time.Time `gorm:"column:latest_time;type:datetime;not null" json:"latest_time"`
	UserID           int64     `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	Passengers       string    `gorm:"column:passengers;type:json" json:"passengers"`
	Vehicles         string    `gorm:"column:vehicles;type:json" json:"vehicles"`
	Status           bool      `gorm:"column:status;type:tinyint(1);not null" json:"status"`
	CreateTime       time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}