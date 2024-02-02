// 定义代码所属的包
package main

//
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// 基本的 web 应用程序
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 Content-Type 类型
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.Header().Set("Content-Type", "/text/plain")
	fmt.Fprint(w, "<h1>Welcometo my great site! </h1>")
}

// 另一个 界面
func cantactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Cantact Page</h1> \n<p>To get in touch: \n<a href=\"https://github.com/baimiyishu13\">github</a>\n</p>\n<p> ")
}

func fqaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> FQA </h1>")
}

// Gin

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/cantact", cantactHandler)
	r.Get("/fqa", fqaHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("404"))
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)
}
