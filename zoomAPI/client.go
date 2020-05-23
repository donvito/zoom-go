package zoomAPI

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url       string
	authToken string
}

func NewClient(apiUrl string, authToken string) (client Client) {

	client.url = apiUrl
	client.authToken = authToken

	return
}

func (client Client) executeRequest(endpoint string, httpMethod string) (response []byte, err error){
	httpClient := &http.Client{}

	url := fmt.Sprintf("%s%s", client.url, endpoint)
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer " + client.authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}


func (client Client) executeRequestWithBody(endpoint string, httpMethod string, body []byte) (response []byte, err error){
	httpClient := &http.Client{}

	url := fmt.Sprintf("%s%s", client.url, endpoint)

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+client.authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}



