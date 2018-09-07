package fetcher

import (
"golang.org/x/text/transform"
"golang.org/x/text/encoding"
"golang.org/x/text/encoding/unicode"
"golang.org/x/net/html/charset"
"log"
"bufio"
_"io"
"net/http"
"io/ioutil"
"fmt"
"time"
"os"
"math/rand"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
    bytes, err := r.Peek(1024)
    if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
    }
    e, _,_ := charset.DetermineEncoding(bytes, "")
    return e
}

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte,error){
	<- rateLimiter
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//resp, err := http.Get(url)

	userAgent := len(UserAgent)
	client := &http.Client{}
    reqest, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(0)
    }
	//log.Printf("Fetcher userAgent: %d", userAgent)
    reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
    reqest.Header.Add("Connection", "keep-alive")
    //reqest.Header.Add("Cookie", "设置cookie")
    //reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
    //reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
    reqest.Header.Add("User-Agent", UserAgent[rand.Intn(userAgent)])
    resp, err := client.Do(reqest)

    if err != nil {
		return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
    }

	bodyReader := bufio.NewReader(resp.Body)
    e := determineEncoding(bodyReader)
    utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
    return ioutil.ReadAll(utf8Reader)
}
