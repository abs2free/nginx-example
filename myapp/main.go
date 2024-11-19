package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed images/*
var images embed.FS

func main() {
	appName := os.Getenv("APP_NAME")
	port := ":3000"

	// Serve embedded images
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的文件路径
		filePath := r.URL.Path[len("/images/"):]

		// 从嵌入的文件系统中打开文件
		data, err := images.ReadFile(filepath.Join("images", filePath))
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// 设置响应头并写入文件内容
		w.Header().Set("Content-Type", "image/jpeg") // 根据需要设置正确的 MIME 类型
		w.Write(data)
	})

	// Serve the index.html file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		log.Printf("Request served by %s", appName)
	})

	log.Printf("%s is listening on port %s", appName, port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
