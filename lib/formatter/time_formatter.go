package Formatter

import "time"

func CurrentTimestampSecond() int64 {
	return time.Now().Unix()
}

func CurrentTimestampMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func CurrentTimestampNanosecond() int64 {
	return time.Now().UnixNano()
}
