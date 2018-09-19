package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const text = `[2018-09-06T14:04:34.396] [INFO] oth - 读取文件流地址： D:/files/sps/twordpdf/GIX888-contract-3968-20180820.docx
[2018-09-06T14:04:34.398] [INFO] oth - 下载地址： ftpfiles/bword/GIX888-contract-3968-20180820.docx
[2018-09-06T14:04:34.400] [INFO] oth - 上传地址： D:/files/sps/twordpdf/GIX888-contract-3968-20180820.docx`

//const logRes  = `(.+)[INFO] oth - 读取文件流地址：(.+)`

func main() {
	logContents, err := ioutil.ReadFile("regex/logfetcher/oth-2018-09-06.log")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(logContents))
	ParseLogContent("2018-09-06T14:03:58.587", string(logContents))
}

func ParseLogContent(key string, contents string) {
	//logRes  := `.+`+key+`[^[*][INFO].+`
	logRes := `[2018-09-06T14:04:34.400][^[*][INFO].+`

	compile := regexp.MustCompile(logRes)

	//findString := res.FindString(text)
	//submatch := compile.FindAllSubmatch(contents, -1)
	findString := compile.FindAllStringSubmatch(contents, -1)
	for _, val := range findString {
		//fmt.Println(val)
		fmt.Println(val[0])
		//fmt.Println(val[1])
	}

}
