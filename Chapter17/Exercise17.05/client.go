package exercise5

import "net/http"

type AwesomeClient struct{
	url string
}

func (a AwesomeClient) GetUsers() (*http.Response, error){
	return http.Get(a.url)
}