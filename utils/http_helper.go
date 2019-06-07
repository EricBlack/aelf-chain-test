package utils

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"strings"
)

type HttpClient struct {
	BaseUrl string
	Headers map[string]string
	Request *HttpRequest.Request
}

func NewClient(url string) (client HttpClient){
	request := HttpRequest.NewRequest()
	request.DisableKeepAlives(false)

	client = HttpClient{
		BaseUrl: url,
		Request: request,
	}

	return
}

func (request HttpClient) SetDefaultHeaders(){
	headers := map[string]string{
		"content-type": "application/json",
		"connection": "close",
	}

	request.SetHeaders(headers)
}

func (request HttpClient) SetHeaders(headers map[string]string) {
	request.Headers = headers
	request.Request.SetHeaders(request.Headers)
}

func (request HttpClient) AddHeaderItem(key, value string) {
	request.Headers[key] = value
	request.Request.SetHeaders(request.Headers)

}

func (request HttpClient) Get(subUrl string, params map[string]interface{}) (response string) {
	response = ""
	url := fmt.Sprintf("%s%s", request.BaseUrl, subUrl)
	resp, err := request.Request.Get(url, params)
	if err != nil {
		return
	}

	body, err := resp.Body()
	if err != nil {
		return
	}
	response = string(body)

	return
}

func (request HttpClient) Post(subUrl string, params map[string]interface{}) (response string) {
	response = ""
	url := fmt.Sprintf("%s/%s", request.BaseUrl, subUrl)
	resp, err := request.Request.Post(url, params)
	if err != nil {
		return
	}

	body, err := resp.Body()
	if err != nil {
		return
	}
	response = string(body)

	return
}

func (request HttpClient) Put(subUrl string, params map[string]interface{}) (response string) {
	response = ""

	return
}

func (request HttpClient) Delete(subUrl string, params map[string]interface{}) (response string) {
	response = ""

	return
}

func GetQueryUrl(url string, params map[string]interface{}) (requestUrl string){
	if params != nil {
		url +="?"
		var urlItems []string
		for key, value := range params{
			switch value.(type) {
			case string:
				urlItems = append(urlItems, fmt.Sprintf("%s=%s", key, value))
			default:
				urlItems = append(urlItems, fmt.Sprintf("%s=%s", key, value.(string)))
			}
		}
		url += strings.Join(urlItems, "&")
	}else{
		requestUrl = url
	}

	return
}
