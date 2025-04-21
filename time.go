package utils

import (
	"context"
	"time"

	"github.com/fufuok/utils/pools/timerpool"
)

const (
	ChinaTimeZone   = "UTC+8"
	ChinaTimeOffset = 8 * 60 * 60
)

// InitChinaLocation 设置全局时区为中国东八区(GMT+8)
func InitChinaLocation() *time.Location {
	loc := time.FixedZone(ChinaTimeZone, ChinaTimeOffset)
	time.Local = loc
	return loc
}

// InitCSTLocation 初始化默认时区为中国东八区(GMT+8)
// 返回值:
// name: "Asia/Shanghai" 或本地时区名称
// loc: 优先尝试解析中国时区, 失败(Windows)后使用本地时区(time.Local)
// cst: 强制偏移的中国时区, !!!注意: 无法使用 time.LoadLocation(cst.String()) 二次加载
// ok: true 表示初始化中国时区成功, false 表示 local 不一定是中国时区
// Deprecated: 对于 公元1年 这种极早日期, 某些系统的时区数据库可能使用当时的 本地平均时间(LMT),
// 结果: +08:05:43 而非标准 +08:00
func InitCSTLocation() (name string, loc *time.Location, cst *time.Location, ok bool) {
	name = "Asia/Shanghai"
	loc, ok = InitLocation(name)
	time.Local = loc
	name = loc.String()
	cst = time.FixedZone(ChinaTimeZone, ChinaTimeOffset)
	return
}

// InitLocation 解析并初始化本地时区
func InitLocation(name string) (*time.Location, bool) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return time.Local, false
	}
	return loc, true
}

// WaitUntilMinute 等待, 直到 m 分钟
func WaitUntilMinute(m int, t ...time.Time) {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
	} else {
		now = time.Now()
	}
	n := m - now.Minute()
	if n < 0 {
		n += 60
	}
	timer := timerpool.New(time.Duration(n) * time.Minute)
	<-timer.C
	timerpool.Release(timer)
}

// WaitNextMinute 下一分钟, 对齐时间, 0 秒
func WaitNextMinute(t ...time.Time) {
	_ = WaitNextMinuteWithTime(t...)
}

// WaitNextMinuteWithTime 下一分钟, 对齐时间, 0 秒
func WaitNextMinuteWithTime(t ...time.Time) (now time.Time) {
	var start time.Time
	if len(t) > 0 {
		start = t[0]
	} else {
		start = time.Now()
	}
	now = BeginOfMinute(start.Add(time.Minute))
	timer := timerpool.New(now.Sub(start))
	<-timer.C
	timerpool.Release(timer)
	return
}

// WaitUntilSecond 等待, 直到 s 秒
func WaitUntilSecond(s int, t ...time.Time) {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
	} else {
		now = time.Now()
	}
	n := s - now.Second()
	if n < 0 {
		n += 60
	}
	timer := timerpool.New(time.Duration(n) * time.Second)
	<-timer.C
	timerpool.Release(timer)
}

// WaitNextSecond 下一秒, 对齐时间, 0 毫秒 (近似)
func WaitNextSecond(t ...time.Time) {
	_ = WaitNextSecondWithTime(t...)
}

// WaitNextSecondWithTime 下一秒, 对齐时间, 0 毫秒 (近似)
func WaitNextSecondWithTime(t ...time.Time) (now time.Time) {
	var start time.Time
	if len(t) > 0 {
		start = t[0]
	} else {
		start = time.Now()
	}
	now = BeginOfSecond(start.Add(time.Second))
	timer := timerpool.New(now.Sub(start))
	<-timer.C
	timerpool.Release(timer)
	return
}

// BeginOfDay 当天 0 点
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay 当天最后时刻
func EndOfDay(t time.Time) time.Time {
	return BeginOfTomorrow(t).Add(-time.Nanosecond)
}

// BeginOfYesterday 昨天 0 点
func BeginOfYesterday(t time.Time) time.Time {
	return BeginOfDay(t.AddDate(0, 0, -1))
}

// EndOfYesterday 昨天最后时刻
func EndOfYesterday(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, -1))
}

// BeginOfTomorrow 明天 0 点
func BeginOfTomorrow(t time.Time) time.Time {
	return BeginOfDay(t.AddDate(0, 0, 1))
}

// EndOfTomorrow 明天 0 点
func EndOfTomorrow(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, 1))
}

// BeginOfSecond 0 毫秒
func BeginOfSecond(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

// EndOfSecond 最后一毫秒
func EndOfSecond(t time.Time) time.Time {
	return BeginOfSecond(t).Add(time.Second - time.Nanosecond)
}

// BeginOfMinute 0 秒
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute 最后一秒
func EndOfMinute(t time.Time) time.Time {
	return BeginOfMinute(t).Add(time.Minute - time.Nanosecond)
}

// BeginOfHour 0 分
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 最后一分
func EndOfHour(t time.Time) time.Time {
	return BeginOfHour(t).Add(time.Hour - time.Nanosecond)
}

// BeginOfWeek 本周一 0 点
func BeginOfWeek(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return BeginOfDay(t).AddDate(0, 0, offset)
}

// EndOfWeek 本周末最后一刻
func EndOfWeek(t time.Time) time.Time {
	return BeginOfNextWeek(t).Add(-time.Nanosecond)
}

// BeginOfLastWeek 上周一 0 点
func BeginOfLastWeek(t time.Time) time.Time {
	return BeginOfWeek(t.AddDate(0, 0, -7))
}

// EndOfLastWeek 上周一最后一刻
func EndOfLastWeek(t time.Time) time.Time {
	return EndOfWeek(t.AddDate(0, 0, -7))
}

// BeginOfNextWeek 下周一 0 点
func BeginOfNextWeek(t time.Time) time.Time {
	return BeginOfWeek(t.AddDate(0, 0, 7))
}

// EndOfNextWeek 下周一最后一刻
func EndOfNextWeek(t time.Time) time.Time {
	return EndOfWeek(t.AddDate(0, 0, 7))
}

// BeginOfMonth 当月第一天 0 点
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 当月最后一刻
func EndOfMonth(t time.Time) time.Time {
	return BeginOfNextMonth(t).Add(-time.Nanosecond)
}

// BeginOfLastMonth 上月第一天 0 点
func BeginOfLastMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, -1))
}

// EndOfLastMonth 上月最后一刻
func EndOfLastMonth(t time.Time) time.Time {
	return BeginOfMonth(t).Add(-time.Nanosecond)
}

// BeginOfNextMonth 下月第一天 0 点
func BeginOfNextMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, 31))
}

// EndOfNextMonth 下月最后一刻
func EndOfNextMonth(t time.Time) time.Time {
	return BeginOfMonth(BeginOfMonth(t).AddDate(0, 0, 62)).Add(-time.Nanosecond)
}

// GetMonthDays 当月天数
func GetMonthDays(t time.Time) int {
	return int(BeginOfNextMonth(t).Sub(BeginOfMonth(t)).Hours() / 24)
}

// BeginOfYear 本年第一天 0 点
func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 本年最后一刻
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Sleep 支持上下文中断的 time.Sleep
func Sleep(ctx context.Context, interval time.Duration) error {
	timer := timerpool.New(interval)
	select {
	case <-ctx.Done():
		timerpool.Release(timer)
		return ctx.Err()
	case <-timer.C:
		timerpool.Release(timer)
		return nil
	}
}

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInYear 返回年份天数
func DaysInYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

var daysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysInMonth 返回月份天数
func DaysInMonth(year int, m time.Month) int {
	if m == time.February && IsLeapYear(year) {
		return 29
	}
	return daysInMonth[m]
}
