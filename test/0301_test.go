package test

import (
	"HaibaraAi123/gospider/spider"
	"fmt"
	"regexp"
	"testing"
)

func Test_0301(t *testing.T) {
	URL := "https://chenfeng123.cn"
	content := spider.GetURL(URL)
	re := regexp.MustCompile("<*")
	fmt.Println(re.FindString(string(content)))
}
