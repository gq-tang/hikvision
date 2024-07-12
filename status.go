/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/12 14:20
*/

package hikvision

import (
	"context"
	"net/http"
	"time"
)

type HistoryStatusReq struct {
	IndexCode    string `json:"indexCode"`           // 资源编码
	StartDate    string `json:"startDate,omitempty"` // 开始日期，yyyy-MM-dd格式
	EndDate      string `json:"endDate,omitempty"`   // 结束日期，yyyy-MM-dd格式
	PageNo       int    `json:"pageNo,omitempty"`    // 页码,从1开始
	PageSize     int    `json:"pageSize,omitempty"`  // 页大小,正整数，最大值1000，默认值500
	ResourceType string `json:"resourceType"`        // 资源类型
}

type (
	HistoryStatusItem struct {
		CollectTime time.Time `json:"collectTime"` // 状态采集时间，时间格式为ISO格式
		Online      int       `json:"online"`      // 资源在线状态
	}
	HistoryStatusData struct {
		Total int                 `json:"total"` // 总条数
		List  []HistoryStatusItem `json:"list"`  // 在线记录列表
	}
	HistoryStatusResp struct {
		Code string            `json:"code"` // 返回码
		Msg  string            `json:"msg"`  // 返回描述
		Data HistoryStatusData `json:"data"` // 返回数据
	}
)

// HistoryStatus 获取资源的历史在线记录接口
func (c *Client) HistoryStatus(ctx context.Context, req *HistoryStatusReq) (*HistoryStatusResp, error) {
	var resp HistoryStatusResp
	if err := c.do(ctx, http.MethodPost, PathHistoryStatus, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CameraStatusReq struct {
	RegionId       string   `json:"regionId,omitempty"`       // 区域id；
	IncludeSubNode string   `json:"includeSubNode,omitempty"` // 是否包含下级区域中的资源数据，1包含，0不包含（若regionId为空，则该参数不起作用）
	IndexCodes     []string `json:"indexCodes,omitempty"`     // 监控点编号列表，最大500
	Status         string   `json:"status,omitempty"`         // 状态，1-在线，0-离线，-1-未检测
	PageNo         int      `json:"pageNo,omitempty"`         // 范围 ( 0 , ~ )，不填默认为1
	PageSize       int      `json:"pageSize,omitempty"`       // 范围 ( 0 , 1000 ]，不填默认为10
}

type (
	CameraStatusItem struct {
		DeviceType      string    `json:"deviceType"`      // 设备型号
		DeviceIndexCode string    `json:"deviceIndexCode"` // 设备唯一编码
		RegionIndexCode string    `json:"regionIndexCode"` // 区域编码
		CollectTime     time.Time `json:"collectTime"`     // 采集时间
		RegionName      string    `json:"regionName"`      // 区域名字
		IndexCode       string    `json:"indexCode"`       // 资源唯一编码
		Cn              string    `json:"cn"`              // 设备名称
		TreatyType      string    `json:"treatyType"`      // 码流传输协议，0：UDP，1：TCP
		Manufacturer    string    `json:"manufacturer"`    // 厂商，hikvision-海康，dahua-大华
		Ip              string    `json:"ip"`              // ip地址，监控点无此值
		Port            int       `json:"port"`            // 端口，监控点无此值
		Online          int       `json:"online"`          // 在线状态，0离线，1在线
	}
	CameraStatusData struct {
		PageNo    int                `json:"pageNo"`    // 页码
		PageSize  int                `json:"pageSize"`  // 每页记录数
		TotalPage int                `json:"totalPage"` // 总页数
		Total     int                `json:"total"`     // 总记录数
		List      []CameraStatusItem `json:"list"`      // 资源信息集
	}
	CameraStatusResp struct {
		Code string           `json:"code"` // 返回码，0:接口业务处理成功
		Msg  string           `json:"msg"`  // 返回描述
		Data CameraStatusData `json:"data"` // 数据信息
	}
)

// CameraStatus 获取监控点在线状态
func (c *Client) CameraStatus(ctx context.Context, req *CameraStatusReq) (*CameraStatusResp, error) {
	var resp CameraStatusResp
	if err := c.do(ctx, http.MethodPost, PathCameraStatus, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
