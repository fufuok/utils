<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ntp

```go
import "github.com/fufuok/utils/ntp"
```

Package ntp provides an implementation of a Simple NTP \(SNTP\) client capable of querying the current time from a remote NTP server.  See RFC5905 \(https://tools.ietf.org/html/rfc5905\) for more details.

This approach grew out of a go\-nuts post by Michael Hofmann: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/FlcdMU5fkLQ

## Index

- [Constants](<#constants>)
- [func ClockOffsetChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Duration](<#func-clockoffsetchan>)
- [func Time(host string) (time.Time, error)](<#func-time>)
- [func TimeChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Time](<#func-timechan>)
- [func TimeV(host string, version int) (time.Time, error)](<#func-timev>)
- [type HostResponse](<#type-hostresponse>)
  - [func HostPreferred(hosts []string) *HostResponse](<#func-hostpreferred>)
- [type LeapIndicator](<#type-leapindicator>)
- [type QueryOptions](<#type-queryoptions>)
- [type Response](<#type-response>)
  - [func GetResponse(host string) *Response](<#func-getresponse>)
  - [func Query(host string) (*Response, error)](<#func-query>)
  - [func QueryWithOptions(host string, opt QueryOptions) (*Response, error)](<#func-querywithoptions>)
  - [func (r *Response) Validate() error](<#func-response-validate>)


## Constants

```go
const (
    // LeapNoWarning indicates no impending leap second.
    LeapNoWarning LeapIndicator = 0

    // LeapAddSecond indicates the last minute of the day has 61 seconds.
    LeapAddSecond = 1

    // LeapDelSecond indicates the last minute of the day has 59 seconds.
    LeapDelSecond = 2

    // LeapNotInSync indicates an unsynchronized leap second.
    LeapNotInSync = 3
)
```

## func ClockOffsetChan

```go
func ClockOffsetChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Duration
```

ClockOffsetChan 启动 Simple NTP \(SNTP\), 周期性获取时钟偏移值

## func Time

```go
func Time(host string) (time.Time, error)
```

Time returns the current time using information from a remote NTP server. It uses version 4 of the NTP protocol. On error, it returns the local system time.

## func TimeChan

```go
func TimeChan(ctx context.Context, interval time.Duration, hosts ...string) chan time.Time
```

TimeChan 启动 Simple NTP \(SNTP\), 周期性获取最新时间

## func TimeV

```go
func TimeV(host string, version int) (time.Time, error)
```

TimeV returns the current time using information from a remote NTP server. On error, it returns the local system time. The version may be 2, 3, or 4.

Deprecated: TimeV is deprecated. Use QueryWithOptions instead.

## type HostResponse

```go
type HostResponse struct {
    Host string
    Resp *Response
}
```

### func HostPreferred

```go
func HostPreferred(hosts []string) *HostResponse
```

HostPreferred 选择最快的 NTP Host

## type LeapIndicator

The LeapIndicator is used to warn if a leap second should be inserted or deleted in the last minute of the current month.

```go
type LeapIndicator uint8
```

## type QueryOptions

QueryOptions contains the list of configurable options that may be used with the QueryWithOptions function.

```go
type QueryOptions struct {
    Timeout      time.Duration // defaults to 5 seconds
    Version      int           // NTP protocol version, defaults to 4
    LocalAddress string        // IP address to use for the client address
    Port         int           // Server port, defaults to 123
    TTL          int           // IP TTL to use, defaults to system default
}
```

## type Response

A Response contains time data, some of which is returned by the NTP server and some of which is calculated by the client.

```go
type Response struct {
    // Time is the transmit time reported by the server just before it
    // responded to the client's NTP query.
    Time time.Time

    // ClockOffset is the estimated offset of the client clock relative to
    // the server. Add this to the client's system clock time to obtain a
    // more accurate time.
    ClockOffset time.Duration

    // RTT is the measured round-trip-time delay estimate between the client
    // and the server.
    RTT time.Duration

    // Precision is the reported precision of the server's clock.
    Precision time.Duration

    // Stratum is the "stratum level" of the server. The smaller the number,
    // the closer the server is to the reference clock. Stratum 1 servers are
    // attached directly to the reference clock. A stratum value of 0
    // indicates the "kiss of death," which typically occurs when the client
    // issues too many requests to the server in a short period of time.
    Stratum uint8

    // ReferenceID is a 32-bit identifier identifying the server or
    // reference clock.
    ReferenceID uint32

    // ReferenceTime is the time when the server's system clock was last
    // set or corrected.
    ReferenceTime time.Time

    // RootDelay is the server's estimated aggregate round-trip-time delay to
    // the stratum 1 server.
    RootDelay time.Duration

    // RootDispersion is the server's estimated maximum measurement error
    // relative to the stratum 1 server.
    RootDispersion time.Duration

    // RootDistance is an estimate of the total synchronization distance
    // between the client and the stratum 1 server.
    RootDistance time.Duration

    // Leap indicates whether a leap second should be added or removed from
    // the current month's last minute.
    Leap LeapIndicator

    // MinError is a lower bound on the error between the client and server
    // clocks. When the client and server are not synchronized to the same
    // clock, the reported timestamps may appear to violate the principle of
    // causality. In other words, the NTP server's response may indicate
    // that a message was received before it was sent. In such cases, the
    // minimum error may be useful.
    MinError time.Duration

    // KissCode is a 4-character string describing the reason for a
    // "kiss of death" response (stratum = 0). For a list of standard kiss
    // codes, see https://tools.ietf.org/html/rfc5905#section-7.4.
    KissCode string

    // Poll is the maximum interval between successive NTP polling messages.
    // It is not relevant for simple NTP clients like this one.
    Poll time.Duration
}
```

### func GetResponse

```go
func GetResponse(host string) *Response
```

GetResponse 获取 NTP 响应, 无效值返回 nil

### func Query

```go
func Query(host string) (*Response, error)
```

Query returns a response from the remote NTP server Host. It contains the time at which the server transmitted the response as well as other useful information about the time and the remote server.

### func QueryWithOptions

```go
func QueryWithOptions(host string, opt QueryOptions) (*Response, error)
```

QueryWithOptions performs the same function as Query but allows for the customization of several query options.

### func \(\*Response\) Validate

```go
func (r *Response) Validate() error
```

Validate checks if the response is valid for the purposes of time synchronization.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
