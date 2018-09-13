package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

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

//定义一个安全的nil parse
func NilParse([]byte) ParseResult {
	return ParseResult{}
}
