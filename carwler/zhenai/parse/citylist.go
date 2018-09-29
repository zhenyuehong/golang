package parse

import (
	"golang/carwler/engine"
	"regexp"
)

//<a href="http://www.zhenai.com/zhenghun/hetian" class="">和田</a>
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc:nil,//这里要进行下一个页面的抓取，这里为了先让他编译通过，暂时设置为nil
			//ParserFunc: engine.NilParse,
			//Parser: engine.NewFuncParser(ParseCity,"ParseCity"),
			ParserFunc: ParseCity,
		})
	}
	return result
}
