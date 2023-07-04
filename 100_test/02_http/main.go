package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func IndexHandlers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello, world")
	log.Println("hello world")
}

func main (){
	//
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		log.Fatal(err)
	} else {
		log.SetOutput(logFile)
	}

	http.HandleFunc("/", IndexHandlers)
	err = http.ListenAndServe("127.0.0.1:8088", nil)
	if err != nil {
		fmt.Printf("listen error:[%v]", err.Error())
	}
}