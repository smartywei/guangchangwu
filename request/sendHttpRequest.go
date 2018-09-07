package bookrequest

import (
	"net/http"
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
	result, _ := http.NewRequest("GET", request.Url, nil)

	for k, v := range request.Headers {
		result.Header.Set(k, v)
	}

	q := result.URL.Query()
	for k, v := range request.Data {
		q.Add(k,v)
	}
	result.URL.RawQuery = q.Encode()

	v, _ := client.Do(result)
	//defer v.Body.Close()
	return v
}