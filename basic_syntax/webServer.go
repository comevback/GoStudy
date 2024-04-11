package main

import (
	"fmt"
	"net/http"
)

// 定义一个处理器函数，用于处理 HTTP 请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func WebServer() {
	// 使用 http.HandleFunc 注册处理器函数
	http.HandleFunc("/", helloHandler)

	// 启动 Web 服务器，监听在 8080 端口
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
