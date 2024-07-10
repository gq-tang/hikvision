/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/10 16:40
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gq-tang/hikvision"
)

func main() {
	cli, err := hikvision.NewClient(&hikvision.ClientOption{
		AppKey:    "29666671",
		AppSecret: "empsl21ds3",
		Host:      "http://www.example.com",
		Log:       nil,
		IsDebug:   true,
	})
	if err != nil {
		panic(err)
	}
	resp, err := cli.EventSubscriptionByEventTypes(context.Background(), &hikvision.EventSubscriptionReq{
		EventTypes: []int{131612},
		EventDest:  "http://www.example.com",
		SubType:    0,
		EventLvl:   nil,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(data))
}
