package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handle) http.Handle

func middleware(next_handler http.Handle) http.Handle{
	return http.HandleFunc(func(res http.ResponseWriter, req *http.Request){
		//hz
		next_handler.ServeHTTP(res, req)
	})
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет"))
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandle)))
	//...

	err := http.ListenAndServe(":8080", mux);
	if(err!=nil) {
		panic(err)
	}
} 