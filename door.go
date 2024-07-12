/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/12 13:48
*/

package hikvision

import (
	"context"
	"net/http"
)

type (
	DoorResourceItem struct {
		IndexCode       string `json:"indexCode"`       // 资源唯一标识
		Name            string `json:"name"`            // 资源名称
		ResourceType    string `json:"resourceType"`    // 资源类型
		DoorNo          int    `json:"doorNo"`          // 资源编号
		Description     string `json:"description"`     // 描述
		ParentIndexCode string `json:"parentIndexCode"` // 父级资源编号
		RegionIndexCode string `json:"regionIndexCode"` // 所属区域唯一标识
		RegionPath      string `json:"regionPath"`      // 所属区域路径，由唯一标示组成，最大10级，格式： @根节点@子区域1@子区域2@
		ChannelType     string `json:"channelType"`     // 通道类型
		ChannelNo       string `json:"channelNo"`       // 通道号
		InstallLocation string `json:"installLocation"` // 安装位置
		CapabilitySet   string `json:"capabilitySet"`   // 设备能力集
		ControlOneId    string `json:"controlOneId"`    // 一级控制器id
		ControlTwoId    string `json:"controlTwoId"`    // 二级控制器id
		ReaderInId      string `json:"readerInId"`      // 读卡器1
		ReaderOutId     string `json:"readerOutId"`     // 读卡器2
		DoorSerial      int    `json:"doorSerial"`      // 门序号
		CreateTime      string `json:"createTime"`      // 创建时间
		UpdateTime      string `json:"updateTime"`      // 更新时间
	}
	DoorResourceData struct {
		Total    int                `json:"total"`    // 记录总数
		PageNo   int                `json:"pageNo"`   // 当前页码
		PageSize int                `json:"pageSize"` // 分页大小
		List     []DoorResourceItem `json:"list"`     // 返回数据
	}
	DoorResourceResp struct {
		Code string             `json:"code"` // 返回码 0: 成功
		Msg  string             `json:"msg"`  // 返回描述
		Data DeviceResourceData `json:"data"` // 资源数据列表
	}
)

// DoorResources 获取门禁资源列表v2
func (c *Client) DoorResources(ctx context.Context, _req *NoTypeResourceReq) (*CameraResourceResp, error) {
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
