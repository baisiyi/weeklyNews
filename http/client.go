package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var Client http.Client = clients()

func clients() http.Client {
	return http.Client{
		Timeout: time.Duration(5) * time.Second, //超时时间
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,   //单个路由最大空闲连接数
			MaxConnsPerHost:       100, //单个路由最大连接数
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func Get(url string, args map[string]interface{}, header map[string]interface{}) (reply []byte, err error) {
	res, err := Response(url, args, header)
	defer res.Body.Close()
	if err != nil {
		return
	}
	reply, err = ioutil.ReadAll(res.Body) // 读取响应 body, 返回为 []byte
	if err != nil {
		return
	}
	return
}

func Response(url string, args map[string]interface{}, header map[string]interface{}) (*http.Response, error) {
	if args != nil {
		url = url + "?" + ToValues(args)
	}
	fmt.Println(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	SetHead(request, header)
	res, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SetHead(request *http.Request, header map[string]interface{}) *http.Request {

	for k, v := range header {
		request.Header.Add(k, fmt.Sprintf("%v", v))
	}

	return request
}

func ToValues(args map[string]interface{}) string {
	params := url.Values{}
	if args != nil {
		for k, v := range args {
			params.Set(k, fmt.Sprintf("%v", v))
		}
		return params.Encode()
	}
	return ""
}
