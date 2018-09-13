package bookrequest

import (
	"net/http"
	"time"
	"fmt"
)

type Request struct {
	Url     string
	Headers map[string]string
	Data    map[string]string
}

func (request *Request) Get() *http.Response {

	if len(request.Url) <= 0 {
		panic("url不能为空")
	}

	client := &http.Client{}
	result, err := http.NewRequest("GET", request.Url, nil)

	if err != nil{
		fmt.Println("new request 异常，正在重试 ...")
		time.Sleep(time.Second * 2 )
		return request.Get()
	}

	for k, v := range request.Headers {
		result.Header.Set(k, v)
	}

	q := result.URL.Query()
	for k, v := range request.Data {
		q.Add(k,v)
	}
	result.URL.RawQuery = q.Encode()

	v, err := client.Do(result)

	if err != nil || v.StatusCode != 200 {
		fmt.Println("抓取内容失败，正在重试 ...")
		time.Sleep(time.Second * 2 )
		return request.Get()
	}

	return v
}