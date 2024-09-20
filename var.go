package hikvision

var ignoreHeaderKey = map[string]struct{}{
	"X-Ca-Signature":         {},
	"X-Ca-Signature-Headers": {},
	"Accept":                 {},
	"Content-MD5":            {},
	"Content-Type":           {},
	"Date":                   {},
	"Content-Length":         {},
	"Server":                 {},
	"Connection":             {},
	"Host":                   {},
	"Transfer-Encoding":      {},
	"X-Application-Context":  {},
	"Content-Encoding":       {},
}

/*
【必选】X-Ca-Key：appKey。
【必选】X-Ca-Signature：签名。
【必选】X-Ca-Signature-Headers：参与headers签名计算的header的key转换为小写字母，按照字典排序后多个key之间使用英文逗号分割，组成字符串。
【可选】X-Ca-Timestamp：API 调用者传递时间戳，值为当前时间的毫秒数，即从1970年1月1日起至今的时间转换为毫秒。
【可选】X-Ca-Nonce：API 调用者生成的 UUID，结合时间戳防重放。
*/
const (
	SysHeaderCaKey         = "X-Ca-Key"
	SysHeaderCaSign        = "X-Ca-Signature"
	SysHeaderCaSignHeaders = "X-Ca-Signature-Headers"
	SysHeaderCaTimestamp   = "X-Ca-Timestamp"
	SysHeaderCaNonce       = "X-Ca-Nonce"
	SysHeaderContentMD5    = "Content-MD5"

	HeaderContentType = "Content-Type"
	HeaderAccept      = "Accept"
)

const (
	PathEventSubscriptionByEventTypes = "/artemis/api/eventService/v1/eventSubscriptionByEventTypes"
	PathResourcesByParams             = "/artemis/api/irds/v2/resource/resourcesByParams"
	PathDeviceResource                = "/artemis/api/irds/v2/deviceResource/resources"
	PathHistoryStatus                 = "/artemis/api/nms/v1/online/history_status"
	PathCameraStatus                  = "/artemis/api/nms/v1/online/camera/get"
)

const (
	EventRegionEntrance = 131586 // 进入区域
	EventRegionExiting  = 131587 // 离开区域
)

const (
	ResourceCamera = "camera" // 监控点
	ResourceDoor   = "door"   // 门禁
)
