package ntp

import (
	"context"
	"sort"
	"time"

	"github.com/fufuok/utils/sched"
	"github.com/fufuok/utils/xsync"
)

const (
	// 特殊值: 无效的 ClockOffset 响应
	invalidClockOffset = 123456 * time.Hour

	// 缺省的请求间隔, 每 3 次返回 1 次结果
	defaultInterval = 2 * time.Hour

	// 最少请求间隔
	defaultMinInterval = 20 * time.Second
)

// 缺省的 NTP Host
var defaultNTPHosts = []string{
	"ntp.aliyun.com",
	"ntp.tencent.com",
	"time.cloudflare.com",
	"time.windows.com",
	"time.apple.com",
	"pool.ntp.org",
}

type HostResponse struct {
	Host string
	Resp *Response
}

// ClockOffsetChan 启动 Simple NTP (SNTP), 周期性获取时钟偏移值
func ClockOffsetChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Duration {
	if interval == 0 {
		interval = defaultInterval
	} else if interval < defaultMinInterval {
		interval = defaultMinInterval
	}
	var offsets []int
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
			offset := invalidClockOffset
			if host == "" {
				hs := HostPreferred(hosts)
				if hs != nil {
					host = hs.Host
					offset = hs.Resp.ClockOffset
				}
			} else {
				if resp := GetResponse(host); resp != nil {
					offset = resp.ClockOffset
				}
			}
			if offset != invalidClockOffset {
				offsets = append(offsets, int(offset))
				if len(offsets) == 3 {
					// 去头尾, 取中间值
					sort.Ints(offsets)
					ch <- time.Duration(offsets[1])
					offsets = offsets[:0]
				}
			}
			<-ticker.C
		}
	}()
	return ch
}

// TimeChan 启动 Simple NTP (SNTP), 周期性获取最新时间
func TimeChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Time {
	if interval == 0 {
		interval = defaultInterval
	} else if interval < defaultMinInterval {
		interval = defaultMinInterval
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
