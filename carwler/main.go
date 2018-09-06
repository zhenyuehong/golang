package main

import (
	"golang/carwler/engine"
	"golang/carwler/zhenai/parse"
)

//提取珍爱网 城市和链接
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})
}
