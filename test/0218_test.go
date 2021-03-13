package test

import (
	"HaibaraAi123/gospider/spider"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func Test0218(t *testing.T) {
	/*
		HistoryFilePath := "history.txt"
		f, err := os.OpenFile(HistoryFilePath, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		//在defer 前check err
	*/
	db, err := spider.InitMysqlClient()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	testGroup := []string{
		/*
			{"https://d1.xia12345.com/d/131/2018/09/hdxMBa.mp4"},
			{"https://d1.xia12345.com/video/202011/5fa8dbcd1d290d0b0846d9e8/720.mp4"},
			{"https://d2.xia12345.com/down/90/2018/06/tSrm7MXu.mp4"},
			{"https://d2.xia12345.com/down/93/2019/01/GkVVdvJA.mp4"},
			{"https://d2.xia12345.com/down/2017/4/gc25001/17425234.mp4"},
			{"https://d2.xia12345.com/down/201610/2001/161018088.mp4"},
			{"https://www.4hur88.com/vod/html3/19092_play_1.html"},
		*/
		//{"https://d2.xia12345.com/down/89/2018/09/wmymj74d.mp4"}，
		//{"https://d2.xia12345.com/down/89/2019/03-1/VyPtk4fv.mp4"},
		//{"https://d2.xia12345.com/down/90/2018/05/cQTnEdWQ.mp4"},
		//{""}
		"https://d2.xia12345.com/down/87/2019/12/ke7UYMNn.mp4",
		"https://d2.xia12345.com/down/88/2019/01/qc7uE3bt.mp4",
		"https://d2.xia12345.com/down/87/2019/03-1/JnGWXWQf.mp4",
		"https://d1.xia12345.com/dl2/videos/202002/imhMIo15/downloads/imhMIo15.mp4",
	}
	for _, v := range testGroup {
		tmp := strings.Split(v, "/")
		length := len(tmp)
		tmpMovieUrl := spider.MovieUrl{}
		if err := db.Where("Name=?", tmp[length-1]).Find(&tmpMovieUrl).Error; err == nil {
			log.Println("Same File:", tmpMovieUrl.Url)
			continue
		}
		err := spider.DownLoadFile(tmp[length-1], v)
		if err != nil {
			panic(err)
		}
		/*
			n, _ := f.Seek(0, os.SEEK_END)
			f.WriteAt(append([]byte(v), []byte("\n")...), n)
		*/
		db.Create(&spider.MovieUrl{Name: tmp[length-1], Url: v, CreatAt: time.Now()})
		log.Println("DownLoaded File ", v)
	}
}
func Test021801(t *testing.T) {
	//go print()
	/*
		for i := 0; i != 10; i = i + 1 {
			fmt.Fprintf(os.Stdout, "result is %d\r", i)
			time.Sleep(time.Second * 1)
		}
		fmt.Println("Over")
	*/
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Fprintf(os.Stdout, "\r%.0f%%[%s]", float64(i)/49*100, h)
	}

}
func print() {
	for {
		fmt.Println("\rtest" + strings.Repeat(" ", 50))
		time.Sleep(time.Second)
	}
}
func Test021802(t *testing.T) {
	Current, Total := 1, 199
	rate := float64(Current) / float64(Total)
	i := int(rate * 50.0)
	h := strings.Repeat("=", i) + strings.Repeat(" ", 50-i)
	fmt.Printf("\r%.2f%%[%s]", rate*100, h)
}
func TestRedis0218(t *testing.T) {
	err := spider.InitRedisClient()
	if err != nil {
		log.Println(err)
	}
}
func TestMysql0218(t *testing.T) {
	db, err := spider.InitMysqlClient()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	filename := "history.txt"
	f, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	cnt := 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		tmp := strings.Split(line, "/")
		length := len(tmp)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		if tmp[length-1] == "720.mp4" {
			db.Create(&spider.MovieUrl{Name: strings.Join([]string{fmt.Sprint(cnt), tmp[length-1]}, ""), Url: line, CreatAt: time.Now()})
			cnt++
			continue
		}
		db.Create(&spider.MovieUrl{Name: tmp[length-1], Url: line, CreatAt: time.Now()})

	}
}
func TestMysql021801(t *testing.T) {
	db, err := spider.InitMysqlClient()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	newMovieUrl := spider.MovieUrl{"17826211.mp4", "https://d2.xia12345.com/down/2017/8/27001/17826211.mp4", time.Now()}
	tmpMovieUrl := spider.MovieUrl{}
	if err := db.Where("Name=?", "111").Find(&tmpMovieUrl).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(tmpMovieUrl)
	if err := db.Where("Name=?", newMovieUrl.Name).Find(&tmpMovieUrl).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(tmpMovieUrl)

}
