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
	"flag"
	"fmt"
	"github.com/gq-tang/hikvision"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "http://127.0.0.1:8080", "host url,example(http://127.0.0.1:8080)")
	flag.Parse()

	cli, err := hikvision.NewClient(&hikvision.ClientOption{
		AppKey:    "20072953",
		AppSecret: "CcPBOy805Utmhffdsl1V",
		Host:      host,
		Log:       nil,
		IsDebug:   true,
	})
	if err != nil {
		panic(err)
	}
	//eventSubscript(cli)
	deviceResource(cli)
	//historyStatus(cli)
	//cameraStatus(cli)
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
		EventDest:  "http://192.168.7.244:30000/eventCb",
		SubType:    1,
		EventLvl:   []int{2},
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
	resp, err := cli.CameraResources(context.Background(), &hikvision.NoTypeResourceReq{
		PageNo:   1,
		PageSize: 100,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(data))
}

func historyStatus(cli *hikvision.Client) {
	defer separate("historyStatus")()
	resp, err := cli.HistoryStatus(context.Background(), &hikvision.HistoryStatusReq{
		IndexCode:    "ce91c758-5af4-4539-845a-1b603746c55",
		ResourceType: hikvision.ResourceDoor,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(data))
}

func cameraStatus(cli *hikvision.Client) {
	defer separate("cameraStatus")()
	resp, err := cli.CameraStatus(context.Background(), &hikvision.CameraStatusReq{
		RegionId:       "root000000",
		IncludeSubNode: "1",
		IndexCodes:     []string{"f98y8c28y85y7y213c082yu95yu"},
		Status:         "1",
		PageNo:         1,
		PageSize:       20,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(data))
}
