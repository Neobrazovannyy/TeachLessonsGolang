package main

import (
	"fmt"
	"net/http"
)

// Response - Ответ
func mainPage(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("main page"))
}

func apiPage(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("api page"))
}

func GetHandle(res http.ResponseWriter, req *http.Request){
	if(req.Method != http.MethodGet){
		http.Error(res, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("GET requests"))
}

func ParseQueryParameters(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)

	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	body += "Query parameters ===============\r\n"
	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	res.Write([]byte(body))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/get_handle", GetHandle)
	mux.HandleFunc("/parse_http_query", ParseQueryParameters)

	err := http.ListenAndServe(":8080", mux);
	if(err!=nil) {
		panic(err)
	}
}