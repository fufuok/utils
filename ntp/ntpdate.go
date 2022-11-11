package ntp

import (
	"context"
	"time"

	"github.com/fufuok/utils/sched"
	"github.com/fufuok/utils/xsync"
)

var (
	// 缺省的 NTP Host
	defaultNTPHosts = []string{
		"ntp.aliyun.com",
		"ntp.tencent.com",
		"time.cloudflare.com",
		"time.windows.com",
		"time.apple.com",
		"pool.ntp.org",
	}

	// 缺省的请求间隔
	defaultInterval = 2 * time.Minute

	// 最少请求间隔
	defaultMinInterval = 10 * time.Second
)

type HostResponse struct {
	Host string
	Resp *Response
}

// ClockOffsetChan 启动 Simple NTP (SNTP), 周期性获取时钟偏移值
func ClockOffsetChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Duration {
	if interval < defaultMinInterval {
		interval = defaultInterval
	}
	ch := make(chan time.Duration, 1)
	go func() {
		host := ""
		ticker := time.NewTicker(interval)
		defer func() {
			ticker.Stop()
			close(ch)
		}()
		for {
			select {
			default:
			case <-ctx.Done():
				return
			}
			if host == "" {
				hs := HostPreferred(hosts)
				if hs != nil {
					host = hs.Host
					ch <- hs.Resp.ClockOffset
				}
			} else {
				if resp := GetResponse(host); resp != nil {
					ch <- resp.ClockOffset
				}
			}
			<-ticker.C
		}
	}()
	return ch
}

// TimeChan 启动 Simple NTP (SNTP), 周期性获取最新时间
func TimeChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Time {
	if interval < time.Second {
		interval = defaultInterval
	}
	ch := make(chan time.Time, 1)
	go func() {
		host := ""
		ticker := time.NewTicker(interval)
		defer func() {
			ticker.Stop()
			close(ch)
		}()
		for {
			select {
			default:
			case <-ctx.Done():
				return
			}
			if host == "" {
				hs := HostPreferred(hosts)
				if hs != nil {
					host = hs.Host
					ch <- hs.Resp.Time
				}
			} else {
				if resp := GetResponse(host); resp != nil {
					ch <- resp.Time
				}
			}
			<-ticker.C
		}
	}()
	return ch
}

// HostPreferred 选择最快的 NTP Host
func HostPreferred(hosts []string) *HostResponse {
	if len(hosts) == 0 {
		hosts = defaultNTPHosts
	}

	m := xsync.NewMap()
	s := sched.New()
	for _, host := range hosts {
		s.Add(1)
		s.RunWithArgs(func(v ...interface{}) {
			host := v[0].(string)
			if resp := GetResponse(host); resp != nil {
				m.Store(host, resp)
			}
		}, host)
	}
	s.WaitAndRelease()

	if m.Size() == 0 {
		return nil
	}

	var optimal *HostResponse
	m.Range(func(host string, v interface{}) bool {
		resp := v.(*Response)
		if optimal == nil || optimal.Resp.RTT > resp.RTT {
			optimal = &HostResponse{
				Host: host,
				Resp: resp,
			}
		}
		return true
	})
	return optimal
}

// GetResponse 获取 NTP 响应, 无效值返回 nil
func GetResponse(host string) *Response {
	resp, err := Query(host)
	if err == nil && resp.Validate() == nil {
		return resp
	}
	return nil
}
