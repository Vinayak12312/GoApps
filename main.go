package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(http.ResponseWriter,*http.Request){
		log.Println("yo")
	})
	http.HandleFunc("/yo",func(http.ResponseWriter,*http.Request){
		log.Println("yoyo")
	})
	http.ListenAndServe(":9090",nil)
}