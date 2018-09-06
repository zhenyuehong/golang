package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//定义一个安全的nil parse

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
