package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

//负责从网上抓取一些数据

//为了防止获取速度过快，被对方网站卡住，我们设置rate limiter 限定一下
var rateLimiter = time.Tick(10 * time.Millisecond) //100ms
//这里多个任务一起抢rateLimiter
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	//resp, err := http.Get(url)
	//resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()

	//防止403 err
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wronf status code: %d", resp.StatusCode)
	}

	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	newReader := bufio.NewReader(resp.Body)
	encode := DetermineEncoding(newReader)
	utf8Reader := transform.NewReader(newReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

	//打印出来的html内容有乱码，原因是html的charset＝gbk的，为了把这个charset转化为UTF8的，我们借助一下官方提供的包
	//gopm get -g -v golang.org/x/text

	//可以帮助我们自动识别网页编码的库:golang.org/x/net/html
}

//自动获取encoding的方式
func DetermineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		//panic(err)
		//log.Printf("fetch error: %v", err)
		//返回默认的utf8编码
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
