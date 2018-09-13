package main

import (
	"net/http"
	"log"
	"my_book/controllers"
	"fmt"
)

func main() {
	http.HandleFunc("/", controllers.ViewBookList)       // 设置访问的路由
	http.HandleFunc("/catlist", controllers.ViewCatLogList) // 设置访问的路由
	http.HandleFunc("/content", controllers.ViewContentList) // 设置访问的路由
	http.HandleFunc("/search", controllers.ViewSearchList)   // 设置访问的路由

	err := http.ListenAndServe(":80", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("服务器开启成功！")
}
