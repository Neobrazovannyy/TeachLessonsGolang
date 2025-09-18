package main

import(
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("This is main page"))
}


func main() {
	max := http.NewServeMux()
	max.HandleFunc("/main_page", mainPage)

	addr := ":8080"
	fmt.Printf("Server started at address: http://127.0.0.1%s\r\n", addr)
	
	err := http.ListenAndServe(addr, max)
	if(err!=nil){
		panic(err)
	}
}