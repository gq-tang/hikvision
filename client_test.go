/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/10 14:41
*/

package hikvision

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sign(t *testing.T) {
	cli, err := NewClient(&ClientOption{
		AppKey:    "20072953",
		AppSecret: "CcPBOy805Utmhffdsl1V",
		Host:      "https://192.168.7.251:443",
		Log:       nil,
		IsDebug:   true,
	})
	if err != nil {
		t.Error(err)
		return
	}

	header := map[string]string{
		"Accept": "*/*",
		//"Accept-Encoding":  "gzip, deflate, sdch",
		//"Accept-Language":  "zh-CN,zh;q=0.8",
		//"Connection":       "keep-alive",
		//"Content-Length":   "0",
		"Content-Type": "application/json",
		//"Cookie":           "JSESSIONID=D9C4A515CACAC31211D1612039D062B7",
		//"X-Requested-With": "XMLHttpRequest",
	}
	signHeader := map[string]string{
		SysHeaderCaKey:       cli.appKey,
		SysHeaderCaTimestamp: "1726822771272",
		"x-ca-nonce":         "298f056a-cc32-c199-70b0-defe68f7c893",
		//"header-A":           "A",
		//"header-B":           "b",
	}
	request, err := cli.newRequest(context.Background(), "", "/artemis/api/eventService/v1/eventSubscriptionByEventTypes", header, signHeader,
		[]byte(`{
    "eventTypes": [
        131586,
        131587
    ],
    "eventDest": "http://192.168.7.244/eventCb",
    "subType": 1,
    "eventLvl": [
        2
    ]
}`))
	if err != nil {
		t.Error(err)
	}
	sign := request.Header.Get(SysHeaderCaSign)
	expectSign := "PR/zLKYh0dS6VjajFpSwsQS2lnnWMFtDENIH3z8efJc="
	assert.Equal(t, sign, expectSign)
}
