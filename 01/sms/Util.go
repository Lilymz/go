package main

import (
	"math/rand"
	"strconv"
	"time"
)

func genMsgId() string {
	// 获取当前年月日
	now := time.Now()
	format := now.Format("20060102")
	return format + strconv.Itoa(int(rand.Int31n(99999)))
}
