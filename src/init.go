package src

import (
	"fmt"
	"reflect"

	"github.com/imroc/req/v3"
)

// 单例配置文件
var C *Config

// 单例HTTP客户端
var HTTPClient *req.Request

func Init(c *Config) error {
	client := req.C()
	client.SetBaseURL(c.Host)
	baseRes := new(BaseResponse)
	resp, err := client.R().
		SetHeader("INSTANCE-UUID", c.UUID).
		SetHeader("INSTANCE-TOKEN", c.Token).
		SetResult(baseRes).
		SetError(baseRes).
		Post("/instance/ping")
	if err != nil {
		return err
	}
	if resp.IsError() || !baseRes.Result {
		return fmt.Errorf("响应出现错误 : %s", baseRes.Message)
	} else if resp.IsSuccess() || baseRes.Result {
		C = c
		HTTPClient = req.C().SetBaseURL(c.Host).
			R().
			SetHeader("INSTANCE-UUID", c.UUID).
			SetHeader("INSTANCE-TOKEN", c.Token)
		return nil
	}
	return fmt.Errorf("出现未知错误 : %s , 状态码 : %s", baseRes.Message, resp.Status)

}

func checkInit() (*Config, error) {
	if reflect.DeepEqual(C, Config{}) {
		return nil, fmt.Errorf("SDK未初始化")
	}
	return C, nil
}
