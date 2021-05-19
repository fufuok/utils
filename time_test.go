package utils

import (
	"testing"
	"time"
)

func TestGet0DataTime(t *testing.T) {
	now := time.Date(2020, 2, 18, 12, 13, 14, 123456789, time.UTC)
	AssertEqual(t, "2020-02-18T00:00:00Z", Get0Hour(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:00:00Z", Get0Minute(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-18T12:13:00Z", Get0Second(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-17T00:00:00Z", Get0Yesterday(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-19T00:00:00Z", Get0Tomorrow(now).Format(time.RFC3339Nano))

	AssertEqual(t, "2020-02-17T00:00:00Z", Get0Week(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-10T00:00:00Z", Get0LastWeek(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-02-24T00:00:00Z", Get0NextWeek(now).Format(time.RFC3339Nano))

	AssertEqual(t, "2020-02-01T00:00:00Z", Get0Month(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-01-01T00:00:00Z", Get0LastMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, "2020-03-01T00:00:00Z", Get0NextMonth(now).Format(time.RFC3339Nano))
	AssertEqual(t, 29, GetMonthDays(now))
}
