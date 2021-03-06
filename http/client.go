package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://album.zhenai.com/u/1384052131", nil)
	if err != nil {
		panic(err)
	}
	//request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s\n", bytes)

}
