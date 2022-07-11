package utils

import "time"

//GetUTCTime returns utc ms for current time
func GetUTCTime() int64 {
	return int64(time.Now().Unix() * 1000)
}

// GetMillisByTime returns utc ms for given time object
func GetMillisByTime(t time.Time) int64 {
	return int64(t.Unix() * 1000)
}
