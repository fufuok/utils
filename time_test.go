package utils

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/fufuok/utils/assert"
)

func TestGet0DataTime(t *testing.T) {
	now := time.Date(2020, 2, 18, 12, 13, 14, 123456789, time.UTC)
	assert.Equal(t, "2020-02-18T00:00:00Z", BeginOfDay(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T23:59:59.999999999Z", EndOfDay(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:00:00Z", BeginOfHour(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:59:59.999999999Z", EndOfHour(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:13:00Z", BeginOfMinute(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:13:59.999999999Z", EndOfMinute(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:13:14Z", BeginOfSecond(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-18T12:13:14.999999999Z", EndOfSecond(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-17T00:00:00Z", BeginOfYesterday(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-17T23:59:59.999999999Z", EndOfYesterday(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-19T00:00:00Z", BeginOfTomorrow(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-19T23:59:59.999999999Z", EndOfTomorrow(now).Format(time.RFC3339Nano))

	assert.Equal(t, "2020-02-17T00:00:00Z", BeginOfWeek(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-23T23:59:59.999999999Z", EndOfWeek(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-10T00:00:00Z", BeginOfLastWeek(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-16T23:59:59.999999999Z", EndOfLastWeek(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-24T00:00:00Z", BeginOfNextWeek(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-03-01T23:59:59.999999999Z", EndOfNextWeek(now).Format(time.RFC3339Nano))

	assert.Equal(t, "2020-02-01T00:00:00Z", BeginOfMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-02-29T23:59:59.999999999Z", EndOfMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-01-01T00:00:00Z", BeginOfLastMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-01-31T23:59:59.999999999Z", EndOfLastMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-03-01T00:00:00Z", BeginOfNextMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-03-31T23:59:59.999999999Z", EndOfNextMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, 29, GetMonthDays(now))

	assert.Equal(t, "2020-01-01T00:00:00Z", BeginOfYear(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2020-12-31T23:59:59.999999999Z", EndOfYear(now).Format(time.RFC3339Nano))

	now = time.Date(2019, 1, 31, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "2018-12-01T00:00:00Z", BeginOfLastMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2018-12-31T23:59:59.999999999Z", EndOfLastMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2019-02-01T00:00:00Z", BeginOfNextMonth(now).Format(time.RFC3339Nano))
	assert.Equal(t, "2019-02-28T23:59:59.999999999Z", EndOfNextMonth(now).Format(time.RFC3339Nano))
}

func TestWaitNextSecondWithTime(t *testing.T) {
	now := time.Date(2020, 2, 18, 12, 13, 14, 123456789, time.UTC)
	now = WaitNextSecondWithTime(now)
	assert.Equal(t, "2020-02-18T12:13:15", now.Format("2006-01-02T15:04:05"))
}

func TestSleep(t *testing.T) {
	dur := 50 * time.Millisecond
	ctx := context.Background()
	err := Sleep(ctx, dur)
	assert.Equal(t, nil, err)

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = Sleep(ctx, dur)
	assert.Equal(t, true, errors.Is(err, context.Canceled))
}

func TestDaysInYear(t *testing.T) {
	assert.Equal(t, 365, DaysInYear(1900))
	assert.Equal(t, 365, DaysInYear(2021))
	assert.Equal(t, 366, DaysInYear(2000))
	assert.Equal(t, 366, DaysInYear(2020))
}

func TestDaysInMonth(t *testing.T) {
	m := time.Month(2)
	assert.Equal(t, 28, DaysInMonth(1900, m))
	assert.Equal(t, 31, DaysInMonth(2021, 7))
	assert.Equal(t, 29, DaysInMonth(2000, m))
	assert.Equal(t, 29, DaysInMonth(2020, m))
}

func TestInitCSTLocation(t *testing.T) {
	name, loc, cst, _ := InitCSTLocation()
	assert.NotEmpty(t, name)
	assert.NotNil(t, loc)
	assert.NotNil(t, cst)
	assert.Equal(t, time.Local, loc)
	assert.Equal(t, ChinaTimezone, cst.String())

	ts := time.Now().In(cst).Format(time.RFC3339)
	assert.True(t, strings.Contains(ts, "+08:00"))
}

func TestInitChinaLocation(t *testing.T) {
	loc := InitChinaLocation()
	assert.NotNil(t, loc)
	assert.Equal(t, ChinaTimezone, loc.String())

	ts := time.Now().In(loc).Format(time.RFC3339)
	assert.True(t, strings.Contains(ts, "+08:00"))

	tm := time.Date(1, 1, 1, 0, 0, 0, 0, loc)
	assert.Equal(t, "0001-01-01T00:00:00+08:00", tm.Format(time.RFC3339))
}
