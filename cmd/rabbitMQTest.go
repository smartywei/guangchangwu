package main

import (
	"book/queue"
	"encoding/json"
)

func main()  {

	body,_ := json.Marshal(map[string]string{"name":"霸道总裁"})

	queue.PushToQueue("book_test",body)

	queue.CloseConnAndCh()
}
