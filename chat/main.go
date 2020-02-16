package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

//templは１つのテンプレートを返す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func main() {
	//root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	//start webserver
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

//ServerHTTPはHTTPリクエストを処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}
