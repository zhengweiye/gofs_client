package gofs_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

// 创建客户端
var httpClient = &http.Client{
	//请求超时时间
	Timeout: time.Second * 90,
	// 创建连接池
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 90 * time.Second, // 100-continue状态码超时时间
	},
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageData[T any] struct {
	List          []T   `json:"list"`          // 响应内容
	TotalElements int64 `json:"totalElements"` // 总记录数
	TotalPage     int64 `json:"totalPage"`     // 总页数
}

func download(client *Client, methodPath string, param any) (data []byte, err error) {
	var body *bytes.Reader
	if param == nil {
		param = make(map[string]any)
	}
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return
	}
	body = bytes.NewReader(paramBytes)

	url := fmt.Sprintf("%s://%s:%d/%s/%s", client.http, client.ip, client.port, client.contextPath, methodPath)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("appKey", client.appKey)
	request.Header.Set("appSecret", client.appSecret)
	request.Header.Set("env", client.env)
	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 状态码判断
	if resp.StatusCode != 200 {
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = fmt.Errorf("http请求返回的状态为%s", string(data))
		return
	}

	// 读取io
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(data) == 0 {
		err = fmt.Errorf("返回字节数组为空")
		return
	}
	return
}

func upload[T any](client *Client, methodPath string, fileBytes []byte, fileName string, param map[string]string) (data Result[T], err error) {
	// 创建一个新的multipart编写器
	var buffer bytes.Buffer
	w := multipart.NewWriter(&buffer)
	//defer w.Close()//TODO 不能放在这里

	// 添加表单字段
	if len(param) > 0 {
		for key, value := range param {
			w.WriteField(key, value)
		}
	}

	// 添加文件
	fileWrite, err := w.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		return
	}
	fileWrite.Write(fileBytes)

	// 关闭multipart编写器
	if err = w.Close(); err != nil {
		return
	}

	// 创建HTTP请求
	url := fmt.Sprintf("%s://%s:%d/%s/%s", client.http, client.ip, client.port, client.contextPath, methodPath)
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("appKey", client.appKey)
	req.Header.Set("appSecret", client.appSecret)
	req.Header.Set("env", client.env)

	// 发送请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 状态码判断
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http请求返回的状态为%d", resp.StatusCode)
		return
	}

	// 读取io
	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(resBytes) == 0 {
		err = fmt.Errorf("返回字节数组为空")
		return
	}

	// 反序列化
	var m = make(map[string]any)
	err = json.Unmarshal(resBytes, &m)
	if err != nil {
		return
	}
	err = mapstructure.Decode(m, &data)
	if err != nil {
		return
	}
	return
}

func httpPost[T any](client *Client, methodPath string, param any) (data Result[T], err error) {
	var body *bytes.Reader
	if param == nil {
		param = make(map[string]any)
	}
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return
	}
	body = bytes.NewReader(paramBytes)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s://%s:%d/%s/%s", client.http, client.ip, client.port, client.contextPath, methodPath), body)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("appKey", client.appKey)
	request.Header.Set("appSecret", client.appSecret)
	request.Header.Set("env", client.env)
	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 状态码判断
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http请求返回的状态为%d", resp.StatusCode)
		return
	}

	// 读取io
	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(resBytes) == 0 {
		err = fmt.Errorf("返回字节数组为空")
		return
	}

	// 反序列化
	var m = make(map[string]any)
	err = json.Unmarshal(resBytes, &m)
	if err != nil {
		return
	}
	err = mapstructure.Decode(m, &data)
	if err != nil {
		return
	}
	return
}
