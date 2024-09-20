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
		SysHeaderCaTimestamp: "1726821293519",
		"x-ca-nonce":         "73f44cd3-a584-7f11-1f43-a0d7c2d2317c",
		//"header-A":           "A",
		//"header-B":           "b",
	}
	request, err := cli.newRequest(context.Background(), "", "/artemis/api/irds/v2/region/nodesByParams", header, signHeader, []byte(`{
    "resourceType": "region",
    "pageNo": 1,
    "pageSize": 10
}`))
	if err != nil {
		t.Error(err)
	}
	sign := request.Header.Get(SysHeaderCaSign)
	expectSign := "jnlO79lw7RMAhMuuqONzEGUwHIcsHu3dmai/ML3Odpk="
	assert.Equal(t, sign, expectSign)
}
