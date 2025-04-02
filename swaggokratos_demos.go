package swaggokratos_demos

import (
	"regexp"

	"github.com/yyle88/must"
)

func MustGetPortNum(address string) string {
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)\.(\d+):(\d+)$`)
	matches := re.FindStringSubmatch(address)
	must.Len(matches, 6)
	return matches[5] // 第 5 个捕获组是端口
}
