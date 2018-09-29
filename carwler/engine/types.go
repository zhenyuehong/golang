package engine

import "golang/carwler_distributed/config"

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url string
	//ParserFunc ParserFunc
	Parser Parser
}

//type SerializedParser struct {
//	FunctionName string
//	Args         interface{}
//}

//{"ParseCityList",nil}, {"ProfileParser","userName"}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string //方便查看这个人的详细信息
	Type    string
	Id      string //ID的保护和去重
	Payload interface{}
}

////定义一个安全的nil parse
//func NilParse([]byte) ParseResult {
//	return ParseResult{}
//}

//定义一个安全的nil parse
type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.Parse(contents, url)
}

func (f FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) FuncParser {
	return FuncParser{
		parser: p,
		name:   name,
	}
}
