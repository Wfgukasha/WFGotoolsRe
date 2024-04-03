package wftime

import "time"

// TimeNowFastFormat 能够让您更快输出格式化的时间
func TimeNowFastFormat() string {
	ti := time.Now().Format(time.DateTime)
	return ti
}
