package wechat

import (
	"net/http"
	"reflect"
	"time"
)

// Config is config for wechat pay, all fields is required.
type Config struct {
	AppId string
	MchId string
	Cert  CertSuite

	Apiv3Secret string
	opts        options
}

// CertSuite is the suite for api cert.
type CertSuite struct {
	SerialNo       string
	PrivateKeyTxt  string
	PrivateKeyPath string
}

// Option is optional configuration for wechat pay.
type Option func(o *options)

// Transport set transport to http client.
func Transport(transport http.RoundTripper) Option {
	return func(o *options) {
		if transport == nil || reflect.ValueOf(transport).IsNil() {
			return
		}
		o.transport = transport
	}
}

// Timeout set timeout for http client.
func Timeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

// CertRefreshTime set max cert refresh time, default
// value is 12h.
func CertRefreshTime(refreshTime time.Duration) Option {
	return func(o *options) {
		o.refreshTime = refreshTime
	}
}

// Options return the options
func (c *Config) Options() *options {
	return &c.opts
}

type options struct {
	Domain  string
	Schema  string
	CertUrl string

	transport   http.RoundTripper
	timeout     time.Duration
	refreshTime time.Duration
}

func defaultOptions() options {
	return options{
		Schema:      defaultSchema,
		Domain:      defaultDomain,
		CertUrl:     defaultDomain + "/v3/certificates",
		refreshTime: 12 * time.Hour,
	}
}

const defaultSchema = "WECHATPAY2-SHA256-RSA2048"
const defaultDomain = "https://api.mch.weixin.qq.com"
