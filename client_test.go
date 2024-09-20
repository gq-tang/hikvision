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
		Host:      "https://10.41.101.216:443",
		Log:       nil,
		IsDebug:   true,
	})
	if err != nil {
		t.Error(err)
		return
	}

	header := map[string]string{
		"Accept": "application/json",
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
		SysHeaderCaTimestamp: "1726819980081",
		"x-ca-nonce":         "41b2f3e7-82fd-e498-f312-cb964fa52cc3",
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
	expectSign := "1G918IReC+OIYs3I8TPAt4yTSH6RkMas7fQnb51zuOY="
	assert.Equal(t, sign, expectSign)
}
