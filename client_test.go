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
		AppKey:    "29666671",
		AppSecret: "empsl21ds3",
		Host:      "http://www.example.com",
		Log:       nil,
		IsDebug:   true,
	})
	if err != nil {
		t.Error(err)
		return
	}

	header := map[string]string{
		"Accept":           "*/*",
		"Accept-Encoding":  "gzip, deflate, sdch",
		"Accept-Language":  "zh-CN,zh;q=0.8",
		"Connection":       "keep-alive",
		"Content-Length":   "0",
		"Content-Type":     "text/plain;charset=UTF-8",
		"Cookie":           "JSESSIONID=D9C4A515CACAC31211D1612039D062B7",
		"X-Requested-With": "XMLHttpRequest",
	}
	signHeader := map[string]string{
		SysHeaderCaKey:       cli.appKey,
		SysHeaderCaTimestamp: "1479968678000",
		"header-A":           "A",
		"header-B":           "b",
	}
	request, err := cli.newRequest(context.Background(), "", "/artemis/api/example?a-body=a&qa=a&qb=B&x-body=x", header, signHeader, nil)
	if err != nil {
		t.Error(err)
	}
	sign := request.Header.Get(SysHeaderCaSign)
	exceptSign := "JRpUpk1ETjzr5gsbo4qoEA9EiQPejvNz12B837xV5HI="
	assert.Equal(t, sign, exceptSign)
}
