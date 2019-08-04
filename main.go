package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	never:=make(chan bool)
	for i:=1;i<=300;i++{
		go  aa()
	}
	<-never
}

func aa()  {
	resp,err:= http.Get("http://localhost:8080/buy?user=1&good=1&shop=1&number=1")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}