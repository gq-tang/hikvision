/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/11 9:37
*/

package hikvision

import (
	"context"
	"net/http"
)

// NoTypeResourceReq 获取资源列表v2请求参数
type NoTypeResourceReq struct {
	PageNo   int `json:"pageNo"`   // 当前页码
	PageSize int `json:"pageSize"` // 分页大小
}

// DeviceResourceReq 获取资源列表v2请求参数
type DeviceResourceReq struct {
	PageNo       int    `json:"pageNo"`       // 当前页码
	PageSize     int    `json:"pageSize"`     // 分页大小
	ResourceType string `json:"resourceType"` // 资源类型
}

type (
	DeviceResourceData struct {
		Total    int                      `json:"total"`    // 记录总数
		PageNo   int                      `json:"pageNo"`   // 当前页码
		PageSize int                      `json:"pageSize"` // 分页大小
		List     []map[string]interface{} `json:"list"`     // 返回数据
	}
	DeviceResourceResp struct {
		Code string             `json:"code"` // 返回码 0: 成功
		Msg  string             `json:"msg"`  // 返回描述
		Data DeviceResourceData `json:"data"` // 资源数据列表
	}
)

// DeviceResources 获取资源列表v2
func (c *Client) DeviceResources(ctx context.Context, req *DeviceResourceReq) (*DeviceResourceResp, error) {
	var resp DeviceResourceResp
	if err := c.do(ctx, http.MethodPost, PathDeviceResource, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type (
	ResourcesByParamsExpression struct {
		Key      string   `json:"key"`      // 资源属性名，支持按updateTime、createTime、indexCode进行查询，例如：key传updateTime，operator传between可以查询特定时间段更新的数据，考虑到校时和夏令时，建议值查询过去一天的数据变更
		Operator int      `json:"operator"` // 操作运算符， 0 ：= 1 ：>= 2 ：<= 3 ：in 4 ：not in 5 ：between 6 ：like 7 ：pre like 8 ：suffix like
		Values   []string `json:"values"`   // 资源属性值，=、>=、<=、like、values数组长度只能是1； in、not in，values数组长度大于1，最大不超时20； between只能用于整形、日期（ISO8601格式） ；like只能用于字符串。
	}
	ResourcesByParamsReq struct {
		Name             string                        `json:"name,omitempty"`             // 名称，模糊搜索，最大长度32，若包含中文，最大长度指不超过按照指定编码的字节长度，即getBytes(“utf-8”).length
		RegionIndexCodes []string                      `json:"regionIndexCodes,omitempty"` // 区域编号,支持根据区域批量查询
		IsSubRegion      bool                          `json:"isSubRegion"`                // rue时，搜索regionIndexCodes及其子孙区域的资源； false时，只搜索 regionIndexCodes的资源； isSubRegion不为空，regionIndexCodes也不可为空
		ResourceType     string                        `json:"resourceType"`               // 资源类型
		PageNo           int                           `json:"pageNo"`                     // 当前页码
		PageSize         int                           `json:"pageSize"`                   // 分页大小
		AuthCodes        []string                      `json:"authCodes,omitempty"`        // 权限码集合
		CapabilitySet    []string                      `json:"capabilitySet,omitempty"`    // 设备能力集(含设备上的智能能力)
		OrderBy          string                        `json:"orderBy,omitempty"`          // 排序字段,注意：排序字段必须是查询条件，否则返回参数错误
		OrderType        string                        `json:"orderType,omitempty"`        // 降序升序,降序：desc 升序：asc
		Expressions      []ResourcesByParamsExpression `json:"expressions,omitempty"`      // 查询表达式
	}
)

type (
	ResourcesByParamsData struct {
		Total    int                      `json:"total"`    // 记录总数
		PageNo   int                      `json:"pageNo"`   // 当前页码
		PageSize int                      `json:"pageSize"` // 分页大小
		List     []map[string]interface{} `json:"list"`     // 返回数据
	}
	ResourcesByParamsResp struct {
		Code string                `json:"code"` // 返回码 0: 成功
		Msg  string                `json:"msg"`  // 返回描述
		Data ResourcesByParamsData `json:"data"`
	}
)

// ResourcesByParams 查询资源列表v2
func (c *Client) ResourcesByParams(ctx context.Context, req *ResourcesByParamsReq) (*ResourcesByParamsResp, error) {
	var resp ResourcesByParamsResp
	if err := c.do(ctx, http.MethodPost, PathResourcesByParams, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
