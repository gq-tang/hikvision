/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/12 13:41
*/

package hikvision

import (
	"context"
	"net/http"
)

type (
	CameraResourceItem struct {
		IndexCode         string `json:"indexCode"`         // 资源唯一标识
		RegionIndexCode   string `json:"regionIndexCode"`   // 所属区域唯一标识
		RegionPath        string `json:"regionPath"`        // 所属区域路径，由唯一标示组成，最大10级，格式： @根节点@子区域1@子区域2@
		ExternalIndexCode string `json:"externalIndexCode"` // 	监控点国标编号
		Name              string `json:"name"`              // 资源名称
		ParentIndexCode   string `json:"parentIndexCode"`   // 父级资源编号
		Longitude         string `json:"longitude"`         // 经度
		Latitude          string `json:"latitude"`          // 纬度
		Elevation         string `json:"elevation"`         // 海拔高度
		CameraType        int    `json:"cameraType"`        // 监控点类型
		InstallLocation   string `json:"installLocation"`   // 安装位置
		ChanNum           int    `json:"chanNum"`           // 通道号
		CascadeCode       string `json:"cascadeCode"`       // 	级联编号
		DacIndexCode      string `json:"dacIndexCode"`      // 所属DAC编号
		Capability        string `json:"capability"`        // 设备能力集(含设备上的智能能力)
		RecordLocation    string `json:"recordLocation"`    // 录像存储位置
		ChannelType       string `json:"channelType"`       // 通道类型
		TransType         int    `json:"transType"`         // 传输协议
		TreatyType        string `json:"treatyType"`        // 接入协议
		CreateTime        string `json:"createTime"`        // 创建时间
		UpdateTime        string `json:"updateTime"`        // 更新时间
	}
	CameraResourceData struct {
		Total    int                  `json:"total"`    // 记录总数
		PageNo   int                  `json:"pageNo"`   // 当前页码
		PageSize int                  `json:"pageSize"` // 分页大小
		List     []CameraResourceItem `json:"list"`     // 返回数据
	}
	CameraResourceResp struct {
		Code string             `json:"code"` // 返回码 0: 成功
		Msg  string             `json:"msg"`  // 返回描述
		Data CameraResourceData `json:"data"` // 资源数据列表
	}
)

// CameraResources 获取摄像头资源列表v2
func (c *Client) CameraResources(ctx context.Context, _req *NoTypeResourceReq) (*CameraResourceResp, error) {
	var resp CameraResourceResp
	req := DeviceResourceReq{
		PageNo:       _req.PageNo,
		PageSize:     _req.PageSize,
		ResourceType: ResourceCamera,
	}
	if err := c.do(ctx, http.MethodPost, PathDeviceResource, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
