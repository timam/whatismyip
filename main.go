package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://whatismyipapi.herokuapp.com/")
	if err != nil{
		fmt.Println("Something Went Wrong!!", err)
	}else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Fatal(err)
		}else {
			bodyString := string(body)
			fmt.Println("Your public IP address is:", bodyString)
		}
	}
}