/**
@description hikvision文件

@copyright    Copyright 2024
@version      1.0.0
@author       tgq
@datetime     2024/7/10 14:16
*/

package hikvision

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func getSign(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
