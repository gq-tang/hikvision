/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/10 16:11
*/

package hikvision

import (
	"context"
	"net/http"
)

// EventSubscriptionReq 按事件类型订阅事件请求参数
type EventSubscriptionReq struct {
	EventTypes []int  `json:"eventTypes"`         // 事件类型
	EventDest  string `json:"eventDest"`          // 指定事件接收的地址，采用restful回调模式，支持http和https
	SubType    int    `json:"subType,omitempty"`  // 订阅类型，0-订阅原始事件，1-联动事件，2-原始事件和联动事件，不填使用默认值0
	EventLvl   []int  `json:"eventLvl,omitempty"` // 事件等级，0-未配置，1-低，2-中，3-高 此处事件等级是指在事件联动中配置的等级 订阅类型为0时，此参数无效，使用默认值0 在订阅类型为1时，不填使用默认值[1,2,3] 在订阅类型为2时，不填使用默认值[0,1,2,3] 数组大小不超过32，事件等级大小不超过31
}

type CommonResp struct {
	Code string `json:"code"` // 返回码 0表示订阅成功，其他表示失败
	Msg  string `json:"msg"`  // 返回描述-记录接口执行情况说明信息 success表示成功描述，其他表示失败
}

// EventSubscriptionByEventTypes 按事件类型订阅事件
func (c *Client) EventSubscriptionByEventTypes(ctx context.Context, req *EventSubscriptionReq) (*CommonResp, error) {
	var resp CommonResp
	if err := c.do(ctx, http.MethodPost, PathEventSubscriptionByEventTypes, nil, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
