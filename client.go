package hikvision

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Logger interface {
	Debugf(format string, args ...interface{})
}

type defaultLog struct{}

func (defaultLog) Debugf(format string, v ...any) {
	log.Printf(format, v...)
}

// Client 通联客户端
type Client struct {
	cli       *http.Client
	appKey    string
	appSecret string
	host      string
	log       Logger
	isDebug   bool
}

type ClientOption struct {
	AppKey    string
	AppSecret string
	Host      string
	Log       Logger
	IsDebug   bool
}

func (opt *ClientOption) validate() error {
	if opt.AppKey == "" {
		return errors.New("AppKey is empty")
	}
	if opt.AppSecret == "" {
		return errors.New("AppSecret is empty")
	}
	if opt.Host == "" {
		return errors.New("host is empty")
	}

	return nil
}

func NewClient(opt *ClientOption) (*Client, error) {
	if err := opt.validate(); err != nil {
		return nil, err
	}

	httpCli := &http.Client{}
	if strings.HasPrefix(strings.ToLower(opt.Host), "https") {
		tlsCfg := &tls.Config{
			InsecureSkipVerify: true,
		}
		transport := &http.Transport{
			TLSClientConfig: tlsCfg,
			Proxy:           http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		httpCli.Transport = transport
	}
	cli := &Client{
		cli:       httpCli,
		appKey:    opt.AppKey,
		appSecret: opt.AppSecret,
		host:      opt.Host,
		log:       opt.Log,
		isDebug:   opt.IsDebug,
	}
	if cli.log == nil && cli.isDebug {
		cli.log = defaultLog{}
	}
	return cli, nil
}

// HTTP METHOD + "\n" +
// Accept + "\n" +     //建议显示设置 Accept Header，部分 Http 客户端当 Accept 为空时会给 Accept
// 设置默认值：*/*，导致签名校验失败。
// Content-MD5  + "\n" +
// Content-Type + "\n" +
// Date + "\n" +
// Headers +
// Url
func (c *Client) sign(method string, path string, req *http.Request, signHeader map[string]string, data []byte) {
	sb := strings.Builder{}
	sb.WriteString(method)
	sb.WriteByte('\n')
	accept := req.Header.Get(HeaderAccept)
	sb.WriteString(accept)
	sb.WriteByte('\n')

	contentType := req.Header.Get(HeaderContentType)
	sb.WriteString(contentType)
	sb.WriteByte('\n')

	keys := make([]string, 0, len(signHeader))
	for k := range signHeader {
		_, ignored := ignoreHeaderKey[k]
		if !ignored {
			keys = append(keys, k)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	lowerKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		v := signHeader[key]
		lowerKeys = append(lowerKeys, strings.ToLower(key))
		sb.WriteString(fmt.Sprintf("%s:%s", strings.ToLower(key), v))
		sb.WriteByte('\n')
	}
	sb.WriteString(path)
	signString := sb.String()
	sign := getSign(signString, c.appSecret)
	req.Header.Set(SysHeaderCaSign, sign)
	signHeaders := strings.Join(lowerKeys, ",")
	req.Header.Set(SysHeaderCaSignHeaders, signHeaders)
	if c.isDebug {
		c.log.Debugf("signString:\n%s", signString)
		c.log.Debugf("signHeaders:%s", signHeaders)
		c.log.Debugf("sign:%s", sign)
	}
}

func (c *Client) newRequest(ctx context.Context, method string, path string, header map[string]string, signHeader map[string]string, data []byte) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if method == "" {
		method = "POST"
	}

	var (
		body io.Reader
	)
	if len(data) > 0 {
		body = bytes.NewBuffer(data)
	}

	url := fmt.Sprintf("%s%s", c.host, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if len(data) > 0 {
		md := md5.New()
		md.Write(data)
		sum := md.Sum(nil)
		contentMd5 := base64.StdEncoding.EncodeToString(sum)
		req.Header.Set(SysHeaderContentMD5, contentMd5)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	for k, v := range signHeader {
		req.Header.Set(k, v)
	}

	c.sign(method, path, req, signHeader, data)
	return req, nil
}

func (c *Client) doRequest(ctx context.Context, method string, path string, header map[string]string, signHeader map[string]string, req interface{}) (*http.Response, error) {
	var (
		data []byte
		err  error
	)
	if req != nil {
		data, err = json.Marshal(req)
		if err != nil {
			return nil, err
		}
	}
	defaultHeader := map[string]string{
		HeaderContentType: "application/json",
		HeaderAccept:      "*/*",
	}
	for k, v := range header {
		defaultHeader[k] = v
	}

	defaultSignHeader := map[string]string{
		SysHeaderCaKey:       c.appKey,
		SysHeaderCaNonce:     uuid.New().String(),
		SysHeaderCaTimestamp: fmt.Sprintf("%d", time.Now().UnixMilli()),
	}
	for k, v := range signHeader {
		defaultSignHeader[k] = v
	}

	request, err := c.newRequest(ctx, method, path, defaultHeader, defaultSignHeader, data)
	if err != nil {
		return nil, err
	}
	if c.isDebug {
		c.log.Debugf("Header:\n")
		for k, v := range request.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ","))
		}
		c.log.Debugf("Body:\n%s", string(data))
	}
	return c.cli.Do(request)
}

func (c *Client) do(ctx context.Context, method string, path string, header map[string]string, signHeader map[string]string, req interface{}, resp interface{}) error {
	response, err := c.doRequest(ctx, method, path, header, signHeader, req)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New(response.Status)
	}
	contentType := strings.ToLower(response.Header.Get(HeaderContentType))
	if !strings.Contains(contentType, "application/json") {
		return fmt.Errorf("The body is not json")
	}
	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&resp); err != nil {
		return err
	}
	return nil
}
