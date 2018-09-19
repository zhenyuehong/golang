package main

import (
	"golang/carwler/frontend/controller"
	"net/http"
)

func main() {
	//为了解决css js的加载问题
	http.Handle("/", http.FileServer(http.Dir("carwler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("carwler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
