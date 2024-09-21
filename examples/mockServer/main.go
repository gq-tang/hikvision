/**
@description mockServer文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/12 8:50
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gq-tang/hikvision"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		port  int
		isTLS bool
	)
	flag.IntVar(&port, "port", 8080, "web port")
	flag.BoolVar(&isTLS, "tls", false, "TLS enabled")

	flag.Parse()
	http.HandleFunc(hikvision.PathEventSubscriptionByEventTypes, event)
	http.HandleFunc(hikvision.PathDeviceResource, deviceResource)
	http.HandleFunc(hikvision.PathHistoryStatus, historyStatus)
	http.HandleFunc(hikvision.PathCameraStatus, cameraStatus)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("web start at %s\n", address)
	server := &http.Server{Addr: address, Handler: nil}
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-sign:
			server.Close()
		}
	}()

	if isTLS {
		log.Println("start with https")
		if err := server.ListenAndServeTLS("./tls/mycert.crt", "./tls/mykey.pem"); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("server closed")
				return
			}
			fmt.Println("ListenAndServe error: ", err.Error())
		}
	} else {
		log.Println("start with http")
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("server closed")
				return
			}
			fmt.Println("ListenAndServe error: ", err.Error())
		}
	}

}

func separate(fnName string) func() {
	fmt.Printf("--------------%s begin------------\n", fnName)
	return func() {
		fmt.Printf("--------------%s end--------------\n", fnName)
	}
}

func handle(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("content-type", "application/json")
	fmt.Println("header:")
	for k := range r.Header {
		fmt.Printf("%s:%s\n", k, r.Header.Get(k))
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	fmt.Printf("body:\n%s\n", body)
	return nil
}

func event(w http.ResponseWriter, r *http.Request) {
	defer separate(r.RequestURI)()
	if err := handle(w, r); err != nil {
		responseErr(w, -1, err)
		return
	}
	responseSuccess(w)
}

func deviceResource(w http.ResponseWriter, r *http.Request) {
	defer separate(r.RequestURI)()
	if err := handle(w, r); err != nil {
		responseErr(w, -1, err)
		return
	}
	data := `{
    "code": "0",
    "msg": "success",
    "data": {
        "total": 15,
        "pageNo": 1,
        "pageSize": 100,
        "list": [
            {
                "indexCode": "f0ecd30298ed4652baa3d8a5fe7f197e",
                "regionIndexCode": "614b7f14472e43afb08d5cd7e44d24b1",
                "regionPath": "@root000000@614b7f14472e43afb08d5cd7e44d24b1@",
                "externalIndexCode": "",
                "name": "Camera 01",
                "parentIndexCode": "96da0341ec93445291096d8ca2d54c27",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 1,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@vis@io@record@vss@event_io@net@maintenance@event_device@event_vis@status@",
                "recordLocation": "",
                "channelType": "analog",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T17:54:32.138+08:00",
                "updateTime": "2024-09-19T17:54:37.911+08:00"
            },
            {
                "indexCode": "7e3068f4fd494e8bbe7ba6901352f945",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "北侧立杆",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 33,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.775+08:00",
                "updateTime": "2024-09-20T17:33:07.832+08:00"
            },
            {
                "indexCode": "7f273e1af5d246adaeb707cd05b48695",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "一楼走廊",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 34,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.777+08:00",
                "updateTime": "2024-09-20T17:33:10.913+08:00"
            },
            {
                "indexCode": "6f8f248734e946c386348911b7f8ec60",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "西侧云台全景",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 35,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.779+08:00",
                "updateTime": "2024-09-20T17:33:12.635+08:00"
            },
            {
                "indexCode": "fe0dad717ec047e6a41c053788e0ed32",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "大门垛内右侧",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 36,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.781+08:00",
                "updateTime": "2024-09-19T16:50:32.452+08:00"
            },
            {
                "indexCode": "32234e2db9dd4b6cbaec88a0d33c8258",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "东侧立杆",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 37,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.783+08:00",
                "updateTime": "2024-09-20T17:33:14.348+08:00"
            },
            {
                "indexCode": "0853340b7ee14ef88d226aae8d139d3b",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "大门垛内左侧",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 38,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.785+08:00",
                "updateTime": "2024-09-19T16:50:32.454+08:00"
            },
            {
                "indexCode": "0eb99f1b28b6491bb0cbbcee7cb0987c",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "大门垛外右侧",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 39,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.787+08:00",
                "updateTime": "2024-09-19T16:50:32.455+08:00"
            },
            {
                "indexCode": "88299117bafa4c759718485f6fc9d9b9",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "大门垛外左侧",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 40,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.789+08:00",
                "updateTime": "2024-09-19T16:50:32.456+08:00"
            },
            {
                "indexCode": "497264cd848b4f499ca274243afcfb54",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "一楼前门外",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 41,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.791+08:00",
                "updateTime": "2024-09-19T16:50:32.457+08:00"
            },
            {
                "indexCode": "31840868e91f4b63ac0dda52d7b44652",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "南侧立杆",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 42,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.793+08:00",
                "updateTime": "2024-09-20T17:33:16.747+08:00"
            },
            {
                "indexCode": "7286c46a5f1e493a98422115de68fd3c",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "二楼主控室",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 43,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.795+08:00",
                "updateTime": "2024-09-19T16:50:32.459+08:00"
            },
            {
                "indexCode": "2fb45917986744969e0b91b50561644e",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "东侧云台全景",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 44,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.797+08:00",
                "updateTime": "2024-09-20T17:33:18.667+08:00"
            },
            {
                "indexCode": "45536275df9643a6a10dd6e7f52a4ef6",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "西侧全景",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 45,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.799+08:00",
                "updateTime": "2024-09-19T16:50:32.461+08:00"
            },
            {
                "indexCode": "757436f6757e41c08b7a73e7affdefb5",
                "regionIndexCode": "96a633dff26946208cf607d5ee5645a9",
                "regionPath": "@root000000@96a633dff26946208cf607d5ee5645a9@",
                "externalIndexCode": "",
                "name": "东侧全景",
                "parentIndexCode": "6b4dfa4f4a304796903e060866323002",
                "longitude": "",
                "latitude": "",
                "elevation": "",
                "cameraType": 0,
                "installLocation": "",
                "chanNum": 46,
                "cascadeCode": "",
                "dacIndexCode": "--",
                "capability": "@event_face_detect_alarm@event_audio@io@event_face@event_rule@event_veh_compare@event_veh@event_veh_recognition@event_ias@face@event_heat@record@vss@ptz@event_io@net@event_real_time_thermometry@maintenance@event_device@status@",
                "recordLocation": "1",
                "channelType": "digital",
                "transType": 1,
                "treatyType": "hiksdk_net",
                "createTime": "2024-09-19T15:59:56.801+08:00",
                "updateTime": "2024-09-19T16:50:32.462+08:00"
            }
        ]
    }
}`
	w.Write([]byte(data))
}

func historyStatus(w http.ResponseWriter, r *http.Request) {
	defer separate(r.RequestURI)()
	if err := handle(w, r); err != nil {
		responseErr(w, -1, err)
		return
	}
	data := `{
    "code": "0",
    "msg": "Operation succeeded",
    "data": {
        "total": 1,
        "list": [
            {
                "collectTime": "2018-12-28T10:21:40.000+08:00",
                "online": 1
            }
        ]
    }
}`
	w.Write([]byte(data))
}

func cameraStatus(w http.ResponseWriter, r *http.Request) {
	defer separate(r.RequestURI)()
	if err := handle(w, r); err != nil {
		responseErr(w, -1, err)
		return
	}
	data := `{
    "code": "0",
    "msg": "Operation succeeded",
    "data": {
        "pageNo": 1,
        "pageSize": 20,
        "totalPage": 0,
        "total": 1,
        "list": [
            {
                "deviceType": null,
                "regionIndexCode": "root000000",
                "collectTime": "2019-12-04T14:00:02.000+08:00",
                "deviceIndexCode": null,
                "port": null,
                "ip": null,
                "regionName": "根节点",
                "indexCode": "9b04256009bf4260bc6be5f333cbb5e9",
                "online": 1,
                "cn": "IPdome--球机123",
                "treatyType": "1",
                "manufacturer": null
            }
        ]
    }
}
`
	w.Write([]byte(data))
}

type CommonResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func responseErr(w http.ResponseWriter, code int, err error) {
	resp := CommonResp{
		Code: fmt.Sprintf("%d", code),
		Msg:  err.Error(),
	}
	data, _ := json.Marshal(resp)
	_, _ = w.Write(data)
}

func responseSuccess(w http.ResponseWriter) {
	resp := CommonResp{
		Code: "0",
		Msg:  "success",
	}
	data, _ := json.Marshal(resp)
	_, _ = w.Write(data)
}
