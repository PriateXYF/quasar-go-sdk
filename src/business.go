package src

import (
	"fmt"
	"strconv"
)

// 获得业务实例
func Business(name string) (*business, error) {
	_, err := checkInit()
	if err != nil {
		return nil, err
	}
	busiRes := new(BusinessResponse)
	resp, err := HTTPClient.SetFormData(map[string]string{
		"name": name,
	}).SetResult(busiRes).
		Post("/instance/getBusiness")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("响应出现错误 : %s", busiRes.Message)
	}
	// 如果无此业务，则创建默认的业务数据
	if !busiRes.Result && busiRes.Message == "操作失败，失败原因 : 查询业务失败!无此业务" {
		b, err := createBusiness(business{
			Name:         name,
			IsAutoRotate: 1,
			Algorithm:    "aes",
			Note:         "SDK自动生成",
			IsStrict:     0,
		})
		if err != nil {
			return nil, err
		}
		if err := b.LoadKeks(); err != nil {
			return nil, err
		}
		return b, nil
	}
	// 如果有业务但出现错误
	if !busiRes.Result {
		return nil, fmt.Errorf("%s", busiRes.Message)
	}
	b := busiRes.Data
	// 如果是严格模式，直接返回
	if b.IsStrict != 0 {
		return &b, nil
	}
	// 如果不是严格模式，获取其所有密钥，写入缓存
	if err := b.LoadKeks(); err != nil {
		return nil, err
	}
	return &b, nil
}

// 创建业务
func createBusiness(b business) (*business, error) {
	_, err := checkInit()
	if err != nil {
		return nil, err
	}
	busiRes := new(BusinessResponse)
	resp, err := HTTPClient.SetFormData(map[string]string{
		"name":         b.Name,
		"isAutoRotate": strconv.Itoa(b.IsAutoRotate),
		"isStrict":     strconv.Itoa(b.IsStrict),
		"note":         b.Note,
		"algorithm":    b.Algorithm,
	}).SetResult(busiRes).
		Post("/instance/createBusiness")
	if err != nil {
		return nil, err
	}
	if resp.IsError() || !busiRes.Result {
		return nil, fmt.Errorf("响应出现错误 : %s", busiRes.Message)
	}
	return &busiRes.Data, nil
}

// 加载业务的所有kek并写入缓存
func (b *business) LoadKeks() error {
	_, err := checkInit()
	if err != nil {
		return err
	}
	if b.IsStrict != 0 {
		return fmt.Errorf("严格模式下无法获取KEK!")
	}
	keksRes := new(KeksResponse)
	resp, err := HTTPClient.SetFormData(map[string]string{
		"name": b.Name,
	}).SetResult(keksRes).
		Post("/instance/getKeks")
	if err != nil {
		return err
	}
	if resp.IsError() || !keksRes.Result {
		return fmt.Errorf("响应出现错误 : %s", keksRes.Message)
	}
	b.Keks = keksRes.Data
	return nil
}

// 通过版本号获取KEK
func (b business) GetKek(version int) *kek {
	// 如果要求获取最新KEK
	if version == 0 {
		for _, k := range b.Keks {
			if k.ID == b.Kek {
				return &k
			}
		}
		return nil
	}
	// 如果要求获取对应版本KEK
	for _, k := range b.Keks {
		if k.Version == version {
			return &k
		}
	}
	return nil
}
