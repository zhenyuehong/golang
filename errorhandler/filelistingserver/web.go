package main

import (
	"golang/errorhandler/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//统一错误处理
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)

		if err != nil {
			log.Printf("error handling request : %s", err.Error())
			//user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system error
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
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileListing))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type userError interface {
	error            //这个是系统的error，不希望给用户看
	Message() string //这个是可以暴露给用户看的
}
