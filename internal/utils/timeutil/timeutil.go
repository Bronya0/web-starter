package timeutil

import (
	"errors"
	"time"
)

const (
	DefaultLayout     = "2006-01-02 15:04:05"
	DateLayout        = "2006-01-02"
	TimeLayout        = "15:04:05"
	CompactLayout     = "20060102150405"
	RFC3339Layout     = time.RFC3339
	RFC3339NanoLayout = time.RFC3339Nano
)

// ====================== convert ======================

// TimeToTimestamp 将time.Time转换为时间戳（秒或毫秒）
// 如果isMilli为true，返回毫秒时间戳，否则返回秒时间戳
func TimeToTimestamp(t time.Time, isMilli bool) int64 {
	if isMilli {
		return t.UnixNano() / int64(time.Millisecond)
	}
	return t.Unix()
}

// TimestampToTime 将时间戳转换为time.Time
// 自动判断时间戳是秒还是毫秒
func TimestampToTime(timestamp int64) time.Time {
	if timestamp > 1e12 { // 如果时间戳大于1e12，认为是毫秒
		return time.Unix(0, timestamp*int64(time.Millisecond))
	}
	return time.Unix(timestamp, 0)
}

// TimeToStr 将time.Time转换为日期时间字符串，格式为DefaultLayout
func TimeToStr(t time.Time) string {
	return t.Format(DefaultLayout)
}

// StrToTime 将日期时间字符串转换为time.Time
// 格式为DefaultLayout
func StrToTime(datetimeStr string) (time.Time, error) {
	return time.Parse(DefaultLayout, datetimeStr)
}

// TimeToDateStr 将time.Time转换为日期字符串，格式为DateLayout
func TimeToDateStr(t time.Time) string {
	return t.Format(DateLayout)
}

// DateStrToTime 将日期字符串转换为time.Time
// 格式为DateLayout
func DateStrToTime(dateStr string) (time.Time, error) {
	return time.Parse(DateLayout, dateStr)
}

// TimeToTimeStr 将time.Time转换为时间字符串，格式为TimeLayout
func TimeToTimeStr(t time.Time) string {
	return t.Format(TimeLayout)
}

// TimeStrToTime 将时间字符串转换为time.Time
// 格式为TimeLayout
func TimeStrToTime(timeStr string) (time.Time, error) {
	return time.Parse(TimeLayout, timeStr)
}

// TimestampToStr 将时间戳转换为日期时间字符串
// 自动判断时间戳是秒还是毫秒
func TimestampToStr(timestamp int64) string {
	t := TimestampToTime(timestamp)
	return TimeToStr(t)
}

// StrToTimestamp 将日期时间字符串转换为时间戳
// 自动判断返回秒或毫秒时间戳
func StrToTimestamp(datetimeStr string, isMilli bool) (int64, error) {
	t, err := StrToTime(datetimeStr)
	if err != nil {
		return 0, err
	}
	return TimeToTimestamp(t, isMilli), nil
}

// ====================== now ======================

// GetNowTimestamp 获取当前时间的时间戳
// 如果isMilli为true，返回毫秒时间戳，否则返回秒时间戳
func NowTimestamp(isMilli bool) int64 {
	return TimeToTimestamp(time.Now(), isMilli)
}

// NowDateTimeStr 获取当前时间的日期时间字符串
func NowDateTimeStr() string {
	return TimeToStr(time.Now())
}

// NowDateStr 获取当前日期的日期字符串
func NowDateStr() string {
	return TimeToDateStr(time.Now())
}

// NowTimeStr 获取当前时间的时间字符串
func NowTimeStr() string {
	return TimeToTimeStr(time.Now())
}

// ====================== add ======================

// AddSeconds 给指定时间增加秒数
func AddSeconds(t time.Time, seconds int64) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinutes 给指定时间增加分钟数
func AddMinutes(t time.Time, minutes int64) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHours 给指定时间增加小时数
func AddHours(t time.Time, hours int64) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddDays 给指定时间增加天数
func AddDays(t time.Time, days int64) time.Time {
	return t.AddDate(0, 0, int(days))
}

// AddMonths 给指定时间增加月数
func AddMonths(t time.Time, months int64) time.Time {
	return t.AddDate(0, int(months), 0)
}

// AddYears 给指定时间增加年数
func AddYears(t time.Time, years int64) time.Time {
	return t.AddDate(int(years), 0, 0)
}

// ====================== diff ======================

// DiffSeconds 计算两个时间之间的秒数差
func DiffSeconds(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Seconds())
}

// DiffMinutes 计算两个时间之间的分钟数差
func DiffMinutes(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Minutes())
}

// DiffHours 计算两个时间之间的小时数差
func DiffHours(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours())
}

// DiffDays 计算两个时间之间的天数差
func DiffDays(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours() / 24)
}

// ====================== is ======================

// IsLeapYear 判断指定年份是否为闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// ====================== get ======================

// GetDaysInMonth 获取指定年份和月份的天数
func GetDaysInMonth(year int, month time.Month) (int, error) {
	if month < 1 || month > 12 {
		return 0, errors.New("invalid month")
	}
	if month == time.February {
		if IsLeapYear(year) {
			return 29, nil
		}
		return 28, nil
	}
	if month == time.April || month == time.June || month == time.September || month == time.November {
		return 30, nil
	}
	return 31, nil
}

// GetStartOfDay 获取指定时间的当天开始时间（00:00:00）
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取指定时间的当天结束时间（23:59:59）
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// GetStartOfMonth 获取指定时间的当月开始时间
func GetStartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfMonth 获取指定时间的当月结束时间
func GetEndOfMonth(t time.Time) time.Time {
	nextMonth := t.Month() + 1
	if nextMonth > 12 {
		nextMonth = 1
	}
	return time.Date(t.Year(), nextMonth, 0, 23, 59, 59, 0, t.Location())
}

// GetStartOfYear 获取指定时间的当年开始时间
func GetStartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfYear 获取指定时间的当年结束时间
func GetEndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 0, t.Location())
}
