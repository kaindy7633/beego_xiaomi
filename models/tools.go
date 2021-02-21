package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// Md5 加密方法
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

// 获取当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 将时间戳抓换为日期时间
func UnixToDate(timestamp int) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}
