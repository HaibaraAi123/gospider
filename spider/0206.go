package spider

import (
	"fmt"
	"regexp"
)

func RE() {
	str := "chenfeng123.cn"
	re := regexp.MustCompile("chenfeng")
	result := re.FindString(str)
	fmt.Println(result)
}
