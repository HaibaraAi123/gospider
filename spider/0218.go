package spider

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DownLoader io.reader
type DownLoader struct {
	io.Reader
	Total, Current int64
}

func (d *DownLoader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += int64(n)
	rate := float64(d.Current) / float64(d.Total)
	i := int(rate * 50.0)
	h := strings.Repeat("=", i) + strings.Repeat(" ", 50-i)
	fmt.Printf("\r%.2f%%[%s]", rate*100, h)
	return
}

//InitRedisClient 初始化并测试redis
func InitRedisClient() error {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		return err
	}
	fmt.Println(val)

	return nil
}

type MovieUrl struct {
	Name    string `gorm:"unique;not null"`
	Url     string `gorm:"not null"`
	CreatAt time.Time
}

func (u *MovieUrl) TableName() string {
	return "MovieUrl"
}

//InitMysqlClient 测试mysql连接
func InitMysqlClient() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "FengChan:HaibaraAi0715@tcp(127.0.0.1:3306)/chenfeng?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	//db.AutoMigrate(&MovieUrl{})
	return db, nil

}

//DownLoadFile download mp4 file
func DownLoadFile(filepath string, url string) error {
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	resp, err := http.Get(url)
	//_, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return err
	}
	if err != nil {
		panic(err)
	}
	/*
		defer func(){
			resp.Body()
		}
	*/
	defer resp.Body.Close()

	output, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	downLoader := &DownLoader{
		Reader: resp.Body,
		Total:  resp.ContentLength,
	}
	if _, err = io.Copy(output, downLoader); err != nil {
		log.Println(err)
	}

	return err
}

func copyBuffer(dst io.Writer, src io.Reader, buf []byte, length int64) (written int64, err error) {
	/*
		if wt, ok := src.(io.WriterTo); ok {
			return wt.WriteTo(dst)
		}
		if rt, ok := dst.(io.ReaderFrom); ok {
			return rt.ReadFrom(src)
		}
	*/
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}

		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		//FuncCallBack(curent, total)

	}
	return written, err
}
