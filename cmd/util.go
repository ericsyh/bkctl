package cmd

import (
	"strconv"
	"time"
)

func UnixToTime(e int64) (datatime time.Time) {
	raw := strconv.FormatInt(e, 10)
	data, _ := strconv.ParseInt(raw, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return datatime
}
