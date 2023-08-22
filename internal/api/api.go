package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
*@Auther kaikai.zheng
*@Date 2023-08-22 10:35:38
*@Name api
*@Desc // 外部API接口
**/

// 查询旅客票
func ShipTicketList(start, end int64, date string) ([]*TicketData, error) {
	ticketReq := &TicketReq{
		StartPortNo: start,
		EndPortNo:   end,
		StartDate:   date,
	}
	reqBytes, _ := json.Marshal(ticketReq)
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, EnqURL, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	for k, v := range generateHeader() {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 201 {
		return nil, fmt.Errorf("post error status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	resBytes, _ := ioutil.ReadAll(resp.Body)
	var ticketResp *TicketResp
	if err := json.Unmarshal(resBytes, &ticketResp); err != nil {
		return nil, err
	}
	return ticketResp.Data, nil
}

func generateHeader() map[string]string {
	header := map[string]string{
		"Host":            `www.ssky123.com`,
		"Sec-Fetch-Site":  `same-origin`,
		"Accept-Language": `zh-CN,zh-Hans;q=0.9`,
		"Accept-Encoding": `gunzip, deflate, br`,
		"Sec-Fetch-Mode":  `cors`,
		"Origin":          `https://www.ssky123.com`,
		//"authentication":  `1684483703659965`,
		"authentication": `1684807090652251`,
		"User-Agent":     `Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.31(0x18001f37) NetType/WIFI Language/zh_CN`,
		"Referer":        `https://www.ssky123.com/online_booking/`,
		"Content-Length": `0`,
		"Connection":     `keep-alive`,
		"Sec-Fetch-Dest": `empty`,
		"Accept":         `application/json, text/plain, */*`,
		"Content-Type":   "application/json",
	}
	return header
}
