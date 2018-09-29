package worker

import (
	"errors"
	"fmt"
	"golang/carwler/engine"
	"golang/carwler/zhenai/parse"
	"golang/carwler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{

		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parse.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parse.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parse.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}
	default:
		return nil, errors.New("unknown parse name")
	}
}
