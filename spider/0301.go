package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//GetURL get html by url
func GetURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error", resp.StatusCode)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return res
}
