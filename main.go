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

func WhatIsMyIP() string {

	var myIP publicIPAddr
	resp, err := http.Get("https://whatismyip-api.herokuapp.com/")

	if err != nil{
		fmt.Println("Something Went Wrong!!", err)
	}else {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil{
			log.Fatal(err)
		}else {
			textBytes := []byte(body)
			myIP = publicIPAddr{}
			err := json.Unmarshal(textBytes, &myIP)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return myIP.PublicIP
}

func main()  {
	spinner := spinner.StartNew("Waiting for server response...")
	spinner.SetCharset([]string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "})

	myIP := WhatIsMyIP()
	fmt.Printf("\nYour public IP address is: %s\n", myIP)

	spinner.Stop()
}