package wfgtime

import "time"

func TimeNowFastFormat() string {
	ti := time.Now().Format(time.DateTime)
	return ti
}
