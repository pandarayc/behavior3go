package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenUuid() string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机数种子
	hexDigits := "0123456789abcdef"
	s := make([]string, 36)

	for i := 0; i < 36; i++ {
		s[i] = string(hexDigits[rand.Intn(16)])
	}

	// bits 12-15 of the time_hi_and_version field to 0010
	s[14] = "4"

	// bits 6-7 of the clock_seq_hi_and_reserved to 01
	s[19] = string(hexDigits[(s[19][0]&0x3)|0x8])

	// 插入分隔符
	s[8] = "-"
	s[13] = "-"
	s[18] = "-"
	s[23] = "-"

	uuid := strings.Join(s, "")
	return uuid
}
