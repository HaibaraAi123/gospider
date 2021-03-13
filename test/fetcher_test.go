package test

import (
	"HaibaraAi123/gospider/fetcher"
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	url := "https://www.chenfeng123.cn"
	content, err := fetcher.Fetch(url)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	filepath := strings.Split(url, "//")[1]
	file, err := os.Create(strings.Join([]string{filepath, ".html"}, ""))
	defer file.Close()
	if err != nil {
		panic(err)
	}
	Writer := bufio.NewWriter(file)
	Writer.Write(content)
}
