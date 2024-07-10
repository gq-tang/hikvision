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
	"encoding/json"
	"net/http"
)

/*
 */
type EventSubscriptionReq struct {
	EventTypes []int  `json:"eventTypes"`
	EventDest  string `json:"eventDest"`
	SubType    int    `json:"subType,omitempty"`
	EventLvl   []int  `json:"eventLvl"`
}

type EventSubscriptionResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (c *Client) EventSubscriptionByEventTypes(ctx context.Context, req *EventSubscriptionReq) (*EventSubscriptionResp, error) {
	response, err := c.doRequest(ctx, http.MethodPost, PathEventSubscriptionByEventTypes, nil, nil, req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var resp EventSubscriptionResp
	if err := decoder.Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
