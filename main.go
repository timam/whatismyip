package main

import (
	"encoding/json"
	"fmt"
	"github.com/janeczku/go-spinner"
	"io/ioutil"
	"log"
	"net/http"
)

type publicIPAddr struct {
	PublicIP string `json:"public_ip"`
}


func WhatIsMyIP()  {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil{
		fmt.Println("Something Went Wrong!!", err)
	}else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Fatal(err)
		}else {
			textBytes := []byte(body)
			myip := publicIPAddr{}
			err := json.Unmarshal(textBytes, &myip)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(myip.PublicIP)

			//bytes := []byte(body)
			//var IP publicIPAddr
			//mmm := json.Unmarshal(bytes, &IP)
			//fmt.Println(mmm)
		}
	}
}

func main()  {
	s := spinner.StartNew("Waiting for server response...")
	s.SetCharset([]string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "})
	WhatIsMyIP()
	s.Stop()
}