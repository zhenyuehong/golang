package parse

import (
	"golang/carwler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	all := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range all {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc:nil,//这里要进行下一个页面的抓取，这里为了先让他编译通过，暂时设置为nil
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
	}
	return result
}
