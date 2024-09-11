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
    "msg": "SUCCESS",
    "data": {
        "total": 4,
        "pageNo": 1,
        "pageSize": 1,
        "list": [
            {
                "indexCode": "e747cb6f0f3d4762b024a9c8f2e7793f",
                "name": "门b_门1",
                "resourceType": "door",
                "doorNo": 1,
                "description": "null",
                "parentIndexCodes": "aefba9b4208c43f88d1015f951d9e181",
                "regionIndexCode": "0129900000000001",
                "regionPath": "@root000000@",
                "channelType": "door",
                "channelNo": "1",
                "installLocation": "null",
                "capabilitySet": "null",
                "controlOneId": "5b7e23fa-12b7-44be-aad1-f5941b9a53c6",
                "controlTwoId": "null",
                "readerInId": "2aab1eab-d410-45a2-89ac-1409b07d5d7e",
                "readerOutId": "d8a5476e-25c0-4aa2-b7e3-db3788ba1f77",
                "comId": "acs",
                "createTime": "2018-11-28T16:47:27:358+08:00",
                "updateTime": "2018-11-28T16:48:34:011+08:00"
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
