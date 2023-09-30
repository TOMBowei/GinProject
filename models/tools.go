package models

import "time"

// 时间戳转换成日期格式
func TimeToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
