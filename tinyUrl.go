package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

    params := url.Values{}
    params.Add("url", "https://marketplace.digitalocean.com/")
    body := strings.NewReader(params.Encode())

    req, err := http.NewRequest("POST", "https://cleanuri.com/api/v1/shorten", body)

    if err != nil {
        fmt.Println("There was an error with the request", err.Error())
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    
    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        fmt.Println("A response error occured", err.Error())
    }    
    
    respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Fatal(err)
	}
    
    fmt.Println(string(respBody))
    defer resp.Body.Close()
}