package main

import (
	"fmt"
	"github.com/janeczku/go-spinner"
	"io/ioutil"
	"log"
	"net/http"
)

func WhatIsMyIP()  {
	resp, err := http.Get("https://whatismyipapi.herokuapp.com/")
	if err != nil{
		fmt.Println("Something Went Wrong!!", err)
	}else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Fatal(err)
		}else {
			bodyString := string(body)
			fmt.Println("\nYour public IP address is:", bodyString)
		}
	}
}

func main()  {
	s := spinner.StartNew("Waiting for server response...")
	s.SetCharset([]string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "})
	WhatIsMyIP()
	s.Stop()
}