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
	"regexp"
)

//提取珍爱网 城市和链接
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
	//fmt.Printf("%s\n", all)
	//打印出来的html内容有乱码，原因是html的charset＝gbk的，为了把这个charset转化为UTF8的，我们借助一下官方提供的包
	//gopm get -g -v golang.org/x/text

	//可以帮助我们自动识别网页编码的库:golang.org/x/net/html
	printCityList(all)
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

func printCityList(contents []byte) {
	//<a href="http://www.zhenai.com/zhenghun/baoshan"
	//class="">宝山</a>
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	//all := compile.FindAll(contents, -1)
	all := compile.FindAllSubmatch(contents, -1)
	//[][][]byte  ->可以看成[][]string
	for _, m := range all {
		//fmt.Printf("%s\n", m)
		//for _,subMatch := range m{
		//		//	fmt.Printf("%s ", subMatch)
		//		//}
		//		//fmt.Println()
		fmt.Printf("City:%s , URL: %s\n", m[2], m[1])
		//这个m打印出的结果  <a href="http://www.zhenai.com/zhenghun/zunyi"
		//class="">遵义</a>http://www.zhenai.com/zhenghun/zunyi遵义
		//0是全部字符串，1是URL，2是城市
	}

	fmt.Printf("matches found %d\n", len(all))
}
