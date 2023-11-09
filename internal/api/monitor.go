package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/**
*@Auther 
*@Date 2023-05-05 16:42:05
*@Name monitor
*@Desc // 监控告警
**/

var (
	SpiderManQG = "-"
	DateLoyout1 = "2006-01-02 15:04:05"
)

type Msg struct {
	Labels      *Labels       `json:"labels"`
	Annotations []*Annotation `json:"annotations"`
	Users       []*User       `json:"users"`
}

type Labels struct {
	Alertname string `json:"alertname"`
	Facility  string `json:"facility"`
	Job       string `json:"job"`
	Level     string `json:"level"`
	Project   string `json:"project"`
	Extra     string `json:"extra"`
}

type Annotation struct {
	AlertWay int    `json:"alertWay"`
	AlertMsg string `json:"alertMsg"`
	Title    string `json:"title,omitempty"`
	Subject  string `json:"subject,omitempty"`
}

type User struct {
	IdNo  string `json:"idNo"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Memo  string `json:"memo"`
}

// 发送告警消息
func SendMsg(facility, content string) error {
	if len(content) > 1024 {
		content = content[:1024]
	}
	var (
		url             = ""
		ApplicationJson = "application/json"
		currTime        = time.Now().Format(DateLoyout1)
		msgData         = fmt.Sprintf("[%s][%s][%s]\n%s", "", facility, currTime, content)
	)
	msg := &Msg{
		Labels: &Labels{
			Alertname: "airline",
			Facility:  facility,
			Job:       "",
			Level:     "P5",
			Project:   "",
			Extra:     "测试",
		},
		Annotations: []*Annotation{
			{
				AlertWay: 10000,
				AlertMsg: msgData,
			},
			{
				AlertWay: 10101,
				AlertMsg: msgData,
			},
		},
		Users: []*User{
			{
				IdNo:  "",
				Name:  "",
				Email: "",
				Memo:  "",
			},
		},
	}
	reqBytes, _ := json.Marshal(msg)
	data := bytes.NewBuffer(reqBytes)
	resp, err := http.Post(url, ApplicationJson, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("monitor send message failed,err=[%v]", resp)
	}
	return nil
}
