package main

import (
	"example/errorhandler/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//统一错误处理
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("error handling request : %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandlerFileListing))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
