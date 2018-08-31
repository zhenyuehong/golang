package main

import (
	"example/retriever/mock"
	"example/retriever/real"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var r Retriever
	mockRetriever := &mock.Retriever{Contents: "this is a fake test"}
	r = mockRetriever
	//指针接收者的实现只能以指针的方式使用；值接收者两者都可以
	//r = &mock.Retriever{"this is a fake test"}
	fmt.Printf("%T %v", r, r)
	//r = real.Retriever{}
	r = &real.Retriever{"Mozilla/5.0", time.Minute}

	fmt.Printf("%T %v", r, r)
	//fmt.Println(download(r))
	//fmt.Println(download(mock.Retriever{"this is second fake test"}))

	//type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(mockRetriever))
}

//接口的组合
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "hardy",
			"lang": "go",
		})
}

//这时候又有一个
type RestrieverPoster interface {
	Retriever
	Poster
}

func session(poster RestrieverPoster) string {
	poster.Post(url, map[string]string{
		"contents": "another fake baidu",
	})
	return poster.Get(url)
}
