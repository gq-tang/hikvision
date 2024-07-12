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
		Host:      "http://127.0.0.1:8080",
		Log:       nil,
		IsDebug:   false,
	})
	if err != nil {
		panic(err)
	}
	eventSubscript(cli)
	deviceResource(cli)
}

func separate(fnName string) func() {
	fmt.Printf("--------------%s begin------------\n", fnName)
	return func() {
		fmt.Printf("--------------%s end--------------\n", fnName)
	}
}

func eventSubscript(cli *hikvision.Client) {
	defer separate("eventSubscript")()
	resp, err := cli.EventSubscriptionByEventTypes(context.Background(), &hikvision.EventSubscriptionReq{
		EventTypes: []int{hikvision.EventRegionEntrance, hikvision.EventRegionExiting},
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

func deviceResource(cli *hikvision.Client) {
	defer separate("deviceResource")()
	resp, err := cli.DeviceResources(context.Background(), &hikvision.DeviceResourceReq{
		PageNo:       1,
		PageSize:     100,
		ResourceType: hikvision.ResourceCamera,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(data))
}