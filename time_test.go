package utils

import (
	"testing"
	"time"
)

func TestGet0DataTime(t *testing.T) {
	now := time.Date(2020, 2, 18, 12, 13, 14, 123456789, time.UTC)
	AssertEqual(t, "2020-02-18T00:00:00Z", BeginOfDay(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T23:59:59.999999999Z", EndOfDay(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:00:00Z", BeginOfHour(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:59:59.999999999Z", EndOfHour(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:13:00Z", BeginOfMinute(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:13:59.999999999Z", EndOfMinute(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-17T00:00:00Z", BeginOfYesterday(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-17T23:59:59.999999999Z", EndOfYesterday(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-19T00:00:00Z", BeginOfTomorrow(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-19T23:59:59.999999999Z", EndOfTomorrow(now).Format(time.RFC3339Nano))

	AssertEqual(t, "2020-02-17T00:00:00Z", BeginOfWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-23T23:59:59.999999999Z", EndOfWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-10T00:00:00Z", BeginOfLastWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-16T23:59:59.999999999Z", EndOfLastWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-24T00:00:00Z", BeginOfNextWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-03-01T23:59:59.999999999Z", EndOfNextWeek(now).Format(time.RFC3339Nano))

	AssertEqual(t, "2020-02-01T00:00:00Z", BeginOfMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-29T23:59:59.999999999Z", EndOfMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-01-01T00:00:00Z", BeginOfLastMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-01-31T23:59:59.999999999Z", EndOfLastMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-03-01T00:00:00Z", BeginOfNextMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-03-31T23:59:59.999999999Z", EndOfNextMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, 29, GetMonthDays(now))

	AssertEqual(t, "2020-01-01T00:00:00Z", BeginOfYear(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-12-31T23:59:59.999999999Z", EndOfYear(now).Format(time.RFC3339Nano))

	now = time.Date(2019, 1, 31, 0, 0, 0, 0, time.UTC)
	AssertEqual(t, "2018-12-01T00:00:00Z", BeginOfLastMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2018-12-31T23:59:59.999999999Z", EndOfLastMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2019-02-01T00:00:00Z", BeginOfNextMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2019-02-28T23:59:59.999999999Z", EndOfNextMonth(now).Format(time.RFC3339Nano))
}
