package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error status code :", resp.StatusCode)
		return
	}

	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	encode := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	//打印出来的html内容有乱码，原因是html的charset＝gbk的，为了把这个charset转化为UTF8的，我们借助一下官方提供的包
	//gopm get -g -v golang.org/x/text

	//可以帮助我们自动识别网页编码的库:golang.org/x/net/html

}

//自动获取encoding的方式
func determineEncoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
