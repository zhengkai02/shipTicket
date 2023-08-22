package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andybalholm/brotli"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
	resBytes, _ := io.ReadAll(resp.Body)
	var ticketResp *TicketResp
	if err := json.Unmarshal(resBytes, &ticketResp); err != nil {
		return nil, err
	}
	return ticketResp.Data, nil
}

// 车船票查询
func FerryTicketList(start, end int64, date string) ([]*TicketData, error) {
	ticketReq := &TicketReq{
		StartPortNo: start,
		EndPortNo:   end,
		StartDate:   date,
	}
	reqBytes, _ := json.Marshal(ticketReq)
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, FerryEnqURL, bytes.NewReader(reqBytes))
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

// 检查登录状态
func CheckToken(token string) error {
	req, err := http.NewRequest(http.MethodGet, CheckTokenURL, nil)
	if err != nil {
		return err
	}
	for k, v := range generateHeader() {
		req.Header.Set(k, v)
	}
	req.Header.Set("token", token)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	respBytes, err := io.ReadAll(res.Body)
	var checkTokenResp *CheckToeknResp
	if err := json.Unmarshal(respBytes, &checkTokenResp); err != nil {
		return err
	}
	if checkTokenResp.Code != 200 && checkTokenResp.Code != 201 {
		return errors.New(checkTokenResp.Message)
	}
	return nil
}

// 登录
func Login(account, password string) (*LoginResp, error) {
	var (
		client = http.DefaultClient
		header = generateHeader()
	)
	URL := fmt.Sprintf("%s?phoneNum=%s&passwd=%s&deviceType=1", LoginURL, account, password)
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
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
	resBytes, _ := io.ReadAll(resp.Body)
	var loginResp *LoginResp
	if err := json.Unmarshal(resBytes, &loginResp); err != nil {
		return nil, err
	}
	return loginResp, nil
}

// brotli 解压缩
func uncompressed(compressedTxt string) ([]byte, error) {
	b := strings.NewReader(compressedTxt)
	r := brotli.NewReader(b)
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func generateHeader() map[string]string {
	header := map[string]string{
		"accept":          `application/json, text/plain, */*`,
		"sec-fetch-site":  `same-origin`,
		"accept-language": `zh-CN,zh-Hans;q=0.9`,
		"accept-encoding": `ungzip, deflate, br`,
		"sec-fetch-mode":  `cors`,
		"token":           `ssky_user_6f37932ba0494a34a2ecd759b814b399`,
		"authentication":  `1692700248652251`,
		"user-agent":      `Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.31(0x18001f37) NetType/WIFI Language/zh_CN`,
		"referer":         `https://www.ssky123.com/online_booking/`,
		"sec-fetch-dest":  `empty`,
		"Content-Type":    "application/json",
	}
	return header
}
