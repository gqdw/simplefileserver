package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 文件服务器，
type config struct {
	rootpath string
	addr     string
}

var conf config

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&conf)
	if err != nil {
		//panic(err)
		fmt.Println("Error:", err)
	}
	// 这有问题
	//fmt.Println(conf.addr)

}
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.rootpath))
	mux.Handle("/", http.StripPrefix("/", files))
	server := &http.Server{
		//Addr:    "0.0.0.0:8080",
		Addr:    conf.addr,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
