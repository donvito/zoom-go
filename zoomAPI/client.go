package zoomAPI

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

func (client Client) executeRequest(endpoint string, httpMethod string) (response []byte, err error) {
	httpClient := &http.Client{}

	url := fmt.Sprintf("%s%s", client.url, endpoint)
	req, err := http.NewRequest(httpMethod, url, nil)
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

	httpStatusCode := resp.StatusCode
	log.Printf("http response statusCode = %d", httpStatusCode)

	if httpStatusCode != 200 {
		switch httpStatusCode {
		case 401:
			err = errors.New(fmt.Sprintf("unauthorized error %d encountered, " +
				"check your auth token if it's still valid", httpStatusCode))
		default:
			err = errors.New(fmt.Sprintf("http error %d encountered in API call", httpStatusCode))
		}

		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client Client) executeRequestWithBody(endpoint string, httpMethod string, body []byte) (response []byte, err error) {
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

	httpStatusCode := resp.StatusCode
	log.Printf("http response statusCode = %d", httpStatusCode)

	if httpStatusCode != 200 && httpStatusCode != 201 && httpStatusCode != 204 {

		switch httpStatusCode {
		case 401:
			err = errors.New(fmt.Sprintf("unauthorized error %d encountered, " +
				"check your auth token if it's still valid", httpStatusCode))
		default:
			err = errors.New(fmt.Sprintf("http error %d encountered in API call", httpStatusCode))
		}

		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
