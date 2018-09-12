package parse

import (
	"fmt"
	"golang/carwler/engine"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//这里m[2]是拷贝出来的，为了解决m(2) 都只指向一个人的问题
		name := string(m[2])
		fmt.Println()
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc:nil,//这里要进行下一个页面的抓取，这里为了先让他编译通过，暂时设置为nil
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
				//这里m[2] 不是马上运行，而是等到这个循环结束后才排到它，所以在这里用M(2)，早就不是指向这个人了
				//所以要把M(2)拷贝出来 name := string(m[2])
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
